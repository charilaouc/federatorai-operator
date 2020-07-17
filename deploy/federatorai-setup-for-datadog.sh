#!/usr/bin/env bash

show_usage()
{
    cat << __EOF__

    Usage:
        Requirement:
            # Note: -k is for OpenShift env only.
            [-k OpenShift kubeconfig file] # e.g. -k .kubeconfig
                File .kubeconfig can be created by using the following command.
                sh -c "export KUBECONFIG=.kubeconfig; oc login <K8s_LOGIN_URL>"
                e.g. sh -c "export KUBECONFIG=.kubeconfig; oc login https://api.ocp4.example.com:6443"

__EOF__
    exit 1
}

patch_alamedaservice()
{
    alamedaservice_yaml_name="alamedaservice_patch.yaml"
    echo -e "\n$(tput setaf 3)Patching alamedaservice ...$(tput sgr 0)"

    cluster_uid=`kubectl get cm cluster-info -n default -o jsonpath='{.metadata.uid}'`
    if [ "$cluster_uid" = "" ]; then
        echo -e "\n$(tput setaf 1)Error! Failed to get cluster uid.$(tput sgr 0)"
        exit 8
    fi

    kubectl get alamedaservice $alamedaservice_name -n $install_namespace -o jsonpath='{.spec.federatoraiDataAdapter}'|grep -q "CLUSTER_NAME"
    if [ "$?" != "0" ]; then

        cat > $file_folder/${alamedaservice_yaml_name} << __EOF__
spec:
  federatoraiDataAdapter:
    env:
      - name: CLUSTER_NAME
        value: ${cluster_uid}
__EOF__

        kubectl patch alamedaservice $alamedaservice_name -n $install_namespace --type merge --patch "$(cat $file_folder/${alamedaservice_yaml_name})"
        if [ "$?" != "0" ];then
            echo -e "\n$(tput setaf 1)Error! Failed to patch alamedaservice $alamedaservice_name.$(tput sgr 0)"
            exit 8
        fi
        wait_until_pods_ready $max_wait_pods_ready_time 30 $install_namespace 5

    fi
    echo -e "\n$(tput setaf 3)...Done.$(tput sgr 0)"
}

patch_data_adapter_secret()
{
    echo -e "\n$(tput setaf 3)Patching Federator.ai data adapter secret...$(tput sgr 0)"
    if [ "$datadogAPIKey" = "" ] || [ "$datadogApplicationKey" = "" ]; then
        echo -e "\n$(tput setaf 1)Error! API key and Application key can't be empty.$(tput sgr 0)"
        exit
    fi

    secret_name="federatorai-data-adapter-secret"
    secret_yaml_name="adapter-secret.yaml"

    kubectl get secret $secret_name -n $install_namespace -o yaml > $file_folder/$secret_yaml_name
    if [ "$?" != "0" ]; then
        echo -e "\n$(tput setaf 1)Error! Failed to get secret $secret_name$(tput sgr 0)"
        exit 1
    fi

    apikey_base64=`echo -n "$datadogAPIKey" | base64`
    applicationkey_base64=`echo -n "$datadogApplicationKey" | base64`
    sed -i "s|datadog_api_key:.*|datadog_api_key: $apikey_base64|g" $file_folder/$secret_yaml_name
    sed -i "s|datadog_application_key:.*|datadog_application_key: $applicationkey_base64|g" $file_folder/$secret_yaml_name

    kubectl apply -n $install_namespace -f $file_folder/$secret_yaml_name
    if [ "$?" != "0" ]; then
        echo -e "\n$(tput setaf 1)Error! Failed to apply patch on data adapter secret $secret_name$(tput sgr 0)"
        exit 1
    fi
    echo -e "\n$(tput setaf 3)...Done.$(tput sgr 0)"
}

patch_data_adapter_configmap()
{
    if [ "$configure_kafka" != "y" ] && [ "$configure_general" != "y" ]; then
        echo -e "\n$(tput setaf 3)Warning! no kafka nor general application setting will be patching.$(tput sgr 0)"
        return
    fi

    echo -e "\n$(tput setaf 3)Patching Federator.ai data adapter configmap...$(tput sgr 0)"
    configmap_name="federatorai-data-adapter-config"
    configmap_yaml_name="adapter-configmap.yaml"

    if ! curl -sL --fail https://raw.githubusercontent.com/containers-ai/federatorai-operator/${alameda_version}/assets/ConfigMap/federatorai-data-adapter-config.yaml -o $file_folder/$configmap_yaml_name; then
        echo -e "\n$(tput setaf 1)Abort, download federatorai-data-adapter-config.yaml file failed!!!$(tput sgr 0)"
        exit 2
    fi

    sed -i "s|namespace:.*|namespace: $install_namespace|g" $file_folder/$configmap_yaml_name
    # Delete every line after # anchor
    sed -i '/# anchor/,$d' $file_folder/$configmap_yaml_name

    if [ "$configure_general" = "y" ]; then
        # Enable GUI for now, after datadog release. These line must be removed
        sed -i "s|enable_general_dashboard =.*|enable_general_dashboard = true|g" $file_folder/$configmap_yaml_name

        index=0
        count=$(./jq -r '.dataAdapterConfigmapForGeneralApplication | length' $json_file)

        while [[ $index -lt $count ]]; do
            generalApplicationName=$(./jq -r '.dataAdapterConfigmapForGeneralApplication['$index'] | .generalApplicationName' $json_file)
            generalApplicationNamespace=$(./jq -r '.dataAdapterConfigmapForGeneralApplication['$index'] | .generalApplicationNamespace' $json_file)
            minReplicas=$(./jq -r '.dataAdapterConfigmapForGeneralApplication['$index'] | .minReplicas' $json_file)
            maxReplicas=$(./jq -r '.dataAdapterConfigmapForGeneralApplication['$index'] | .maxReplicas' $json_file)

            cat >> $file_folder/$configmap_yaml_name << __EOF__
      [[inputs.datadog]]
        urls = ["\$DATADOG_QUERY_URL"]
        api_key = "\$DATADOG_API_KEY"
        application_key = "\$DATADOG_APPLICATION_KEY"
        # If we keep CLUSTER_NAME value empty, the agent will get k8s cluster name automatically.
        cluster_name = "\$CLUSTER_NAME"
        ## Cloud metric has 5 minutes to 10 minutes delay
        cloud_metric_delay_interval = "30m0s"
        ## Set default cloud information ifneeded
        enable_set_default_cloud_info_if_empty = true
        default_provider = "aws"
        default_region = "us-west1"
        default_instance_type = "m5.4xlarge"
        default_instance_id = "i-00cd730e045190cad"
        default_zone = "us-west-1a"
        # Watched source
        # TOML format reference: https://github.com/influxdata/toml/blob/master/README.md
        [[inputs.datadog.watched_source]]
          namespace = "${generalApplicationNamespace}"
          application = "${generalApplicationName}"
          min_replicas = ${minReplicas} # monitored_application_min_replicas
          max_replicas = ${maxReplicas} # monitored_application_max_replicas
          [[inputs.datadog.watched_source.watched_metrics]]
            name="kubernetes.cpu.usage.total"
            metric_type="CPU_MILLICORES_USAGE"
          [[inputs.datadog.watched_source.watched_metrics]]
            name="kubernetes.memory.usage"
            metric_type="MEMORY_BYTES_USAGE"

      [[inputs.alameda_datahub_query]]
        url = "\$DATAHUB_URL"
        port = "\$DATAHUB_PORT"
        ##The recommendation query range, unit: minutes
        recommendation_interval = 5
        # If we keep CLUSTER_NAME value empty, the agent will get k8s cluster name automatically.
        cluster_name = "\$CLUSTER_NAME"

        [[inputs.alameda_datahub_query.watched_source]] #General Application config
          name = "${generalApplicationName}"
          namespace = "${generalApplicationNamespace}"
          measurement = "controller"
          scope = "prediction"

        [[inputs.alameda_datahub_query.watched_source]] #General Application config
          name = "${generalApplicationName}"
          namespace = "${generalApplicationNamespace}"
          measurement = "controller"
          scope = "recommendation"

        [[inputs.alameda_datahub_query.watched_source]] #General Application config
          name = "${generalApplicationName}"
          namespace = "${generalApplicationNamespace}"
          measurement = "controller"
          scope = "planning"

__EOF__
            ((index = index + 1))
        done
    fi

    if [ "$configure_kafka" = "y" ]; then
        # Enable GUI for now, after datadog release. These line must be removed
        sed -i "s|enable_kafka_dashboard =.*|enable_kafka_dashboard = true|g" $file_folder/$configmap_yaml_name

        index=0
        count=$(./jq -r '.dataAdapterConfigmapForKafka | length' $json_file)

        while [[ $index -lt $count ]]; do
            kafkaConsumerDeploymentName=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerDeploymentName' $json_file)
            kafkaConsumerDeploymentNamespace=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerDeploymentNamespace' $json_file)
            kafkaConsumerMinimumReplica=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerMinimumReplica' $json_file)
            kafkaConsumerMaximumReplica=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerMaximumReplica' $json_file)
            kafkaConsumerGroupName=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerGroupName' $json_file)
            kafkaConsumerGroupNamespace=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerGroupNamespace' $json_file)
            kafkaTopicName=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaTopicName' $json_file)
            kafkaTopicNamespace=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaTopicNamespace' $json_file)

            cat >> $file_folder/$configmap_yaml_name << __EOF__
      [[inputs.datadog_application_aware]]
        urls = ["\$DATADOG_QUERY_URL"]
        api_key = "\$DATADOG_API_KEY"
        application_key = "\$DATADOG_APPLICATION_KEY"
        # If we keep CLUSTER_NAME value empty, the agent will get k8s cluster name automatically.
        cluster_name = "\$CLUSTER_NAME"

        [inputs.datadog_application_aware.watched_kafka_consumer]
          application = "${kafkaConsumerDeploymentName}"
          namespace = "${kafkaConsumerDeploymentNamespace}"
          min_replicas = ${kafkaConsumerMinimumReplica}
          max_replicas = ${kafkaConsumerMaximumReplica}
          topics = ["${kafkaTopicName}"]
          consumer_groups = ["${kafkaConsumerGroupName}"]

      [[inputs.alameda_datahub_query]]
        url = "\$DATAHUB_URL"
        port = "\$DATAHUB_PORT"
        ##The recommendation query range, unit: minutes
        recommendation_interval = 5
        # If we keep CLUSTER_NAME value empty, the agent will get k8s cluster name automatically.
        cluster_name = "\$CLUSTER_NAME"

        [[inputs.alameda_datahub_query.watched_source]]
          name = "${kafkaConsumerGroupName}"
          namespace = "${kafkaConsumerGroupNamespace}"
          measurement = "kafka_consumer_group"
          scope = "recommendation"

        [[inputs.alameda_datahub_query.watched_source]]
          name = "${kafkaConsumerGroupName}"
          namespace = "${kafkaConsumerGroupNamespace}"
          measurement = "kafka_consumer_group_current_offset"
          scope = "prediction"

        [[inputs.alameda_datahub_query.watched_source]]
          name = "${kafkaTopicName}"
          namespace = "${kafkaTopicNamespace}"
          measurement = "kafka_topic_partition_current_offset"
          scope = "prediction"

__EOF__
            ((index = index + 1))
        done
    fi

    kubectl apply -n $install_namespace -f $file_folder/$configmap_yaml_name
    if [ "$?" != "0" ]; then
        echo -e "\n$(tput setaf 1)Error! Failed to apply patch on data adapter configmap.$(tput sgr 0)"
        exit 1
    fi
    echo -e "\n$(tput setaf 3)...Done.$(tput sgr 0)"
}

wait_until_pods_ready()
{
  period="$1"
  interval="$2"
  namespace="$3"
  target_pod_number="$4"

  wait_pod_creating=1
  for ((i=0; i<$period; i+=$interval)); do

    if [[ "$wait_pod_creating" = "1" ]]; then
        # check if pods created
        if [[ "`kubectl get po -n $namespace 2>/dev/null|wc -l`" -ge "$target_pod_number" ]]; then
            wait_pod_creating=0
            echo -e "\nChecking pods..."
        else
            echo "Waiting for pods in namespace $namespace to be created..."
        fi
    else
        # check if pods running
        if pods_ready $namespace; then
            echo -e "\nAll $namespace pods are ready."
            return 0
        fi
        echo "Waiting for pods in namespace $namespace to be ready..."
    fi

    sleep "$interval"

  done

  echo -e "\n$(tput setaf 1)Warning!! Waited for $period seconds, but all pods are not ready yet. Please check $namespace namespace$(tput sgr 0)"
  leave_prog
  exit 4
}

pods_ready()
{
  [[ "$#" == 0 ]] && return 0

  namespace="$1"

  kubectl get pod -n $namespace \
    -o=jsonpath='{range .items[*]}{.metadata.name}{"\t"}{.status.conditions[?(@.type=="Ready")].status}{"\t"}{.status.phase}{"\t"}{.status.reason}{"\n"}{end}' \
      | while read name status phase reason _junk; do
          if [ "$status" != "True" ]; then
            msg="Waiting pod $name in namespace $namespace to be ready."
            [ "$phase" != "" ] && msg="$msg phase: [$phase]"
            [ "$reason" != "" ] && msg="$msg reason: [$reason]"
            echo "$msg"
            return 1
          fi
        done || return 1

  return 0
}

get_datadog_key()
{
    ### datadog info
    echo -e "\nGetting Datadog info..."
    default=$(./jq -r '.dataAdapterSecret.datadogAPIKey' $json_file)
    read -r -p "$(tput setaf 6)Input a Datadog API Key [$default]: $(tput sgr 0)" datadogAPIKey </dev/tty
    datadogAPIKey=${datadogAPIKey:-$default}

    default=$(./jq -r '.dataAdapterSecret.datadogApplicationKey' $json_file)
    read -r -p "$(tput setaf 6)Input a Datadog Application Key [$default]: $(tput sgr 0)" datadogApplicationKey </dev/tty
    datadogApplicationKey=${datadogApplicationKey:-$default}

    configStr=$( ./jq -c $(printf '.dataAdapterSecret.datadogAPIKey="%s"' $datadogAPIKey) <<<$configStr )
    configStr=$( ./jq -c $(printf '.dataAdapterSecret.datadogApplicationKey="%s"' $datadogApplicationKey) <<<$configStr )
    ./jq -r '.' <<< $configStr > $json_file
}

get_kafka_info()
{
    default="y"
    read -r -p "$(tput setaf 3)Do you want to configure configmap for kafka? [default: $default]: $(tput sgr 0): " configure_kafka </dev/tty
    configure_kafka=${configure_kafka:-$default}

    if [ "$configure_kafka" != "y" ]; then
        return
    fi

    index=0
    next_set="y"
    while [[ "$next_set" = "y" ]]
    do
        echo -e "\nGetting Kafka info... No.$((index+1))"
        default=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerDeploymentName' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer deployment name [$default]: $(tput sgr 0)" kafkaConsumerDeploymentName </dev/tty
        kafkaConsumerDeploymentName=${kafkaConsumerDeploymentName:-$default}

        default=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerDeploymentNamespace' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer deployment namespace [$default]: $(tput sgr 0)" kafkaConsumerDeploymentNamespace </dev/tty
        kafkaConsumerDeploymentNamespace=${kafkaConsumerDeploymentNamespace:-$default}

        default=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerMinimumReplica' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer minimum replica number [$default]: $(tput sgr 0)" kafkaConsumerMinimumReplica </dev/tty
        kafkaConsumerMinimumReplica=${kafkaConsumerMinimumReplica:-$default}

        default=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerMaximumReplica' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer maximum replica number [$default]: $(tput sgr 0)" kafkaConsumerMaximumReplica </dev/tty
        kafkaConsumerMaximumReplica=${kafkaConsumerMaximumReplica:-$default}

        default=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerGroupName' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer group name [$default]: $(tput sgr 0)" kafkaConsumerGroupName </dev/tty
        kafkaConsumerGroupName=${kafkaConsumerGroupName:-$default}

        default=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaConsumerGroupNamespace' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer group namespace [$default]: $(tput sgr 0)" kafkaConsumerGroupNamespace </dev/tty
        kafkaConsumerGroupNamespace=${kafkaConsumerGroupNamespace:-$default}

        default=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaTopicName' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer topic name [$default]: $(tput sgr 0)" kafkaTopicName </dev/tty
        kafkaTopicName=${kafkaTopicName:-$default}

        default=$(./jq -r '.dataAdapterConfigmapForKafka['$index'] | .kafkaTopicNamespace' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer topic namespace [$default]: $(tput sgr 0)" kafkaTopicNamespace </dev/tty
        kafkaTopicNamespace=${kafkaTopicNamespace:-$default}

        if [ "$kafkaConsumerDeploymentName" = "" ] || [ "$kafkaConsumerDeploymentNamespace" = "" ] || [ "$kafkaConsumerMinimumReplica" = "" ] || [ "$kafkaConsumerMaximumReplica" = "" ] || [ "$kafkaConsumerGroupName" = "" ] || [ "$kafkaConsumerGroupNamespace" = "" ] || [ "$kafkaTopicName" = "" ] || [ "$kafkaTopicNamespace" = "" ]; then
            echo -e "\n$(tput setaf 1)Error! Kafka info can't be empty.$(tput sgr 0)"
            exit 7
        fi

        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForKafka[%s].kafkaConsumerDeploymentName="%s"' $index $kafkaConsumerDeploymentName) <<<$configStr )
        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForKafka[%s].kafkaConsumerDeploymentNamespace="%s"' $index $kafkaConsumerDeploymentNamespace) <<<$configStr )
        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForKafka[%s].kafkaConsumerMinimumReplica="%s"' $index $kafkaConsumerMinimumReplica) <<<$configStr )
        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForKafka[%s].kafkaConsumerMaximumReplica="%s"' $index $kafkaConsumerMaximumReplica) <<<$configStr )
        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForKafka[%s].kafkaConsumerGroupName="%s"' $index $kafkaConsumerGroupName) <<<$configStr )
        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForKafka[%s].kafkaConsumerGroupNamespace="%s"' $index $kafkaConsumerGroupNamespace) <<<$configStr )
        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForKafka[%s].kafkaTopicName="%s"' $index $kafkaTopicName) <<<$configStr )
        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForKafka[%s].kafkaTopicNamespace="%s"' $index $kafkaTopicNamespace) <<<$configStr )
        ./jq -r '.' <<< $configStr > $json_file
        ((index = index + 1))
        echo ""
        sleep 1

        default="n"
        read -r -p "$(tput setaf 2)Do you want to input another set? [default: n]: $(tput sgr 0)" next_set </dev/tty
        next_set=${next_set:-$default}
        while [[ "$next_set" != "y" ]] && [[ "$next_set" != "n" ]]
        do
            read -r -p "$(tput setaf 2)Do you want to input another set? [default: n]: $(tput sgr 0)" next_set </dev/tty
            next_set=${next_set:-$default}
        done
        if [[ "$next_set" == "n" ]]; then
            configStr=$( ./jq -c $(printf 'del(.dataAdapterConfigmapForKafka[%s:])' $index) <<<$configStr )
            ./jq -r '.' <<< $configStr > $json_file
        fi
    done
}

get_general_application_info()
{
    default="y"
    read -r -p "$(tput setaf 3)Do you want to configure configmap for general application? [default: $default]: $(tput sgr 0): " configure_general </dev/tty
    configure_general=${configure_general:-$default}

    if [ "$configure_general" != "y" ]; then
        return
    fi

    index=0
    next_set="y"
    echo -e "$(tput setaf 127)\nNote!! Use Deployment/DeploymentConfig/Statefulset name as application name.$(tput sgr 0)"
    while [[ "$next_set" = "y" ]]
    do
        echo -e "\nGetting general application info... No.$((index+1))"
        default=$(./jq -r '.dataAdapterConfigmapForGeneralApplication['$index'] | .generalApplicationName' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input application name [$default]: $(tput sgr 0)" generalApplicationName </dev/tty
        generalApplicationName=${generalApplicationName:-$default}

        default=$(./jq -r '.dataAdapterConfigmapForGeneralApplication['$index'] | .generalApplicationNamespace' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input application namespace [$default]: $(tput sgr 0)" generalApplicationNamespace </dev/tty
        generalApplicationNamespace=${generalApplicationNamespace:-$default}

        default=$(./jq -r '.dataAdapterConfigmapForGeneralApplication['$index'] | .minReplicas' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input minimum replicas number [$default]: $(tput sgr 0)" minReplicas </dev/tty
        minReplicas=${minReplicas:-$default}

        default=$(./jq -r '.dataAdapterConfigmapForGeneralApplication['$index'] | .maxReplicas' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input maximum replicas number [$default]: $(tput sgr 0)" maxReplicas </dev/tty
        maxReplicas=${maxReplicas:-$default}

        if [ "$generalApplicationName" = "" ] || [ "$generalApplicationNamespace" = "" ] || [ "$minReplicas" = "" ] || [ "$maxReplicas" = "" ]; then
            echo -e "\n$(tput setaf 1)Error! Application info can't be empty.$(tput sgr 0)"
            exit 7
        fi

        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForGeneralApplication[%s].generalApplicationName="%s"' $index $generalApplicationName) <<<$configStr )
        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForGeneralApplication[%s].generalApplicationNamespace="%s"' $index $generalApplicationNamespace) <<<$configStr )
        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForGeneralApplication[%s].minReplicas="%s"' $index $minReplicas) <<<$configStr )
        configStr=$( ./jq -c $(printf '.dataAdapterConfigmapForGeneralApplication[%s].maxReplicas="%s"' $index $maxReplicas) <<<$configStr )
        ./jq -r '.' <<< $configStr > $json_file
        ((index = index + 1))
        echo ""
        sleep 1

        default="n"
        read -r -p "$(tput setaf 2)Do you want to input another application? [default: n]: $(tput sgr 0)" next_set </dev/tty
        next_set=${next_set:-$default}
        while [[ "$next_set" != "y" ]] && [[ "$next_set" != "n" ]]
        do
            read -r -p "$(tput setaf 2)Do you want to input another application? [default: n]: $(tput sgr 0)" next_set </dev/tty
            next_set=${next_set:-$default}
        done
        if [[ "$next_set" == "n" ]]; then
            configStr=$( ./jq -c $(printf 'del(.dataAdapterConfigmapForGeneralApplication[%s:])' $index) <<<$configStr )
            ./jq -r '.' <<< $configStr > $json_file
        fi
    done
}

get_alamedaservice_full_version()
{
    alameda_version=`kubectl get alamedaservice --all-namespaces|grep -v 'EXECUTION'|awk '{print $4}'`
    if [ "$alameda_version" = "" ]; then
        echo -e "\n$(tput setaf 1)Error, failed to get Federator.ai version!$(tput sgr 0)"
        exit 2
    fi
}

restart_data_adapter_pod()
{
    adapter_pod_name=`kubectl get pods -n $install_namespace -o name |grep "federatorai-data-adapter-"|cut -d '/' -f2`
    if [ "$adapter_pod_name" = "" ]; then
        echo -e "\n$(tput setaf 1)Error, failed to get Federator.ai data adapter pod name!$(tput sgr 0)"
        exit 2
    fi
    kubectl delete pod $adapter_pod_name -n $install_namespace
    if [ "$?" != "0" ]; then
        echo -e "\n$(tput setaf 1)Error! Failed to delete Federator.ai data adapter pod $adapter_pod_name$(tput sgr 0)"
        exit 8
    fi
    wait_until_pods_ready $max_wait_pods_ready_time 30 $install_namespace 5
}

check_version()
{
    openshift_required_minor_version="9"
    k8s_required_version="11"

    oc version 2>/dev/null|grep "oc v"|grep -q " v[4-9]"
    if [ "$?" = "0" ];then
        # oc version is 4-9, passed
        openshift_minor_version="12"
        return 0
    fi

    # OpenShift Container Platform 4.x
    oc version 2>/dev/null|grep -q "Server Version: 4"
    if [ "$?" = "0" ];then
        # oc server version is 4, passed
        openshift_minor_version="12"
        return 0
    fi

    oc version 2>/dev/null|grep "oc v"|grep -q " v[0-2]"
    if [ "$?" = "0" ];then
        # oc version is 0-2, failed
        echo -e "\n$(tput setaf 10)Error! OpenShift version less than 3.$openshift_required_minor_version is not supported by Federator.ai$(tput sgr 0)"
        exit 5
    fi

    # oc major version = 3
    openshift_minor_version=`oc version 2>/dev/null|grep "oc v"|cut -d '.' -f2`
    # k8s version = 1.x
    k8s_version=`kubectl version 2>/dev/null|grep Server|grep -o "Minor:\"[0-9]*.\""|tr ':+"' " "|awk '{print $2}'`

    if [ "$openshift_minor_version" != "" ] && [ "$openshift_minor_version" -lt "$openshift_required_minor_version" ]; then
        echo -e "\n$(tput setaf 10)Error! OpenShift version less than 3.$openshift_required_minor_version is not supported by Federator.ai$(tput sgr 0)"
        exit 5
    elif [ "$openshift_minor_version" = "" ] && [ "$k8s_version" != "" ] && [ "$k8s_version" -lt "$k8s_required_version" ]; then
        echo -e "\n$(tput setaf 10)Error! Kubernetes version less than 1.$k8s_required_version is not supported by Federator.ai$(tput sgr 0)"
        exit 6
    elif [ "$openshift_minor_version" = "" ] && [ "$k8s_version" = "" ]; then
        echo -e "\n$(tput setaf 10)Error! Can't get Kubernetes or OpenShift version$(tput sgr 0)"
        exit 5
    fi
}

prepare_env()
{
    get_alamedaservice_full_version

    if [ -f "$json_file" ]; then
        cat $json_file|grep -q "$json_file_version"
        if [ "$?" != "0" ]; then
            mv $json_file ${json_file}.old
            mv $json_file_template ${json_file_template}.old
        fi
    fi

    if [ ! -f "$json_file" ]; then
        if [ ! -f "$json_file_template" ]; then
            cat > $json_file_template << __EOF__
{
  "version": "${json_file_version}",
  "dataAdapterSecret": {
    "datadogAPIKey": "",
    "datadogApplicationKey": ""
  },
  "dataAdapterConfigmapForKafka": [
    {
      "kafkaConsumerDeploymentName": "",
      "kafkaConsumerDeploymentNamespace": "",
      "kafkaConsumerMinimumReplica": "",
      "kafkaConsumerMaximumReplica": "",
      "kafkaConsumerGroupName": "",
      "kafkaConsumerGroupNamespace": "",
      "kafkaTopicName": "",
      "kafkaTopicNamespace": ""
    }
  ],
  "dataAdapterConfigmapForGeneralApplication": [
    {
      "generalApplicationName": "",
      "generalApplicationNamespace": "",
      "minReplicas": "",
      "maxReplicas": ""
    }
  ]
}
__EOF__
        fi
        cp $json_file_template $json_file
    fi

    if [ ! -f "jq" ]; then
        if ! curl -sL --fail https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 -O; then
            echo -e "\n$(tput setaf 1)Abort, download jq binary failed!!!$(tput sgr 0)"
            exit 2
        fi
        mv jq-linux64 jq
        chmod u+x jq
    fi
    configStr=$(./jq -c '.' $json_file)
}

# If more option is added into script in the future, uncomment this section

# if [ "$#" -eq "0" ]; then
#     show_usage
#     exit 1
# fi

echo "Checking environment version..."
check_version
echo "...Passed"

while getopts "hk:" o; do
    case "${o}" in
        k)
            kubeconfig="${OPTARG}"
            ;;
        h)
            show_usage
            exit 1
            ;;
        *)
            echo "Error! Invalid parameter."
            show_usage
            ;;
    esac
done

[ "$max_wait_pods_ready_time" = "" ] && max_wait_pods_ready_time=1500  # maximum wait time for pods become ready

file_folder="./config_result"
current_location=`pwd`
mkdir -p $file_folder

if [ "$openshift_minor_version" != "" ]; then
    # OpenShift
    if [ "${kubeconfig}" = "" ]; then
        echo -e "\n$(tput setaf 1)Error! Need to use \"-k\" to specify openshift kubeconfig file.$(tput sgr 0)"
        show_usage
    fi
    export KUBECONFIG=${kubeconfig}
fi

# Check if kubectl connect to server.
result="`echo ""|kubectl cluster-info 2>/dev/null`"
if [ "$?" != "0" ]; then
    echo -e "\n$(tput setaf 1)Error! Please login into cluster first.$(tput sgr 0)"
    exit 1
fi
current_server="`echo $result|sed 's/.*at //'|awk '{print $1}'`"
echo "You are connecting to cluster: $current_server"

install_namespace="`kubectl get pods --all-namespaces |grep "alameda-ai-"|awk '{print $1}'|head -1`"
if [ "$install_namespace" = "" ];then
    echo -e "\n$(tput setaf 1)Error! Please install Federatorai before running this script.$(tput sgr 0)"
    exit 3
fi

alamedaservice_name="`kubectl get alamedaservice -n $install_namespace -o jsonpath='{range .items[*]}{.metadata.name}'`"
if [ "$alamedaservice_name" = "" ]; then
    echo -e "\n$(tput setaf 1)Error! Failed to get alamedaservice name.$(tput sgr 0)"
    exit 8
fi

json_file="adapter.json"
json_file_version="v4.3"
json_file_template="adapter.json.tmp"

prepare_env

get_datadog_key

get_general_application_info

get_kafka_info

patch_alamedaservice

patch_data_adapter_secret

patch_data_adapter_configmap

restart_data_adapter_pod

echo -e "$(tput setaf 6)\nSetup Federator.ai for Datadog successfully$(tput sgr 0)"

exit 0
