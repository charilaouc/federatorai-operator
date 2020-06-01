#!/usr/bin/env bash

show_usage()
{
    cat << __EOF__

    Usage:
        Requirement:
            [-k OpenShift kubeconfig file] # e.g. -k .kubeconfig
                File .kubeconfig can be created by using the following command.
                sh -c "export KUBECONFIG=.kubeconfig; oc login <K8s_LOGIN_URL>"
                e.g. sh -c "export KUBECONFIG=.kubeconfig; oc login https://api.ocp4.example.com:6443"

__EOF__
    exit 1
}

patch_data_adapter_secret()
{
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
}

patch_data_adapter_configmap()
{
    configmap_name="federatorai-data-adapter-config"
    configmap_yaml_name="adapter-configmap.yaml"

    if ! curl -sL --fail https://raw.githubusercontent.com/containers-ai/federatorai-operator/${alameda_version}/assets/ConfigMap/federatorai-data-adapter-config.yaml -o $file_folder/$configmap_yaml_name; then
        echo -e "\n$(tput setaf 1)Abort, download federatorai-data-adapter-config.yaml file failed!!!$(tput sgr 0)"
        exit 2
    fi

    sed -i "s|namespace:.*|namespace: $install_namespace|g" $file_folder/$configmap_yaml_name
    # Delete every line after # anchor
    sed -i '/# anchor/,$d' $file_folder/$configmap_yaml_name

    index=0
    count=$(jq -r '.dataAdapterConfigmap | length' $json_file)

    while [[ $index -lt $count ]]; do
        kafkaConsumerDeploymentName=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerDeploymentName' $json_file)
        kafkaConsumerDeploymentNamespace=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerDeploymentNamespace' $json_file)
        kafkaConsumerMinimumReplica=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerMinimumReplica' $json_file)
        kafkaConsumerMaximumReplica=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerMaximumReplica' $json_file)
        kafkaConsumerGroupName=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerGroupName' $json_file)
        kafkaConsumerGroupNamespace=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerGroupNamespace' $json_file)
        kafkaTopicName=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaTopicName' $json_file)
        kafkaTopicNamespace=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaTopicNamespace' $json_file)

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

    kubectl apply -n $install_namespace -f $file_folder/$configmap_yaml_name
    if [ "$?" != "0" ]; then
        echo -e "\n$(tput setaf 1)Error! Failed to apply patch on data adapter configmap.$(tput sgr 0)"
        exit 1
    fi
}

get_datadog_key()
{
    ### datadog info
    echo -e "\nGetting Datadog info..."
    default=$(jq -r '.dataAdapterSecret.datadogAPIKey' $json_file)
    read -r -p "$(tput setaf 6)Input a Datadog API Key [$default]: $(tput sgr 0)" datadogAPIKey </dev/tty
    datadogAPIKey=${datadogAPIKey:-$default}

    default=$(jq -r '.dataAdapterSecret.datadogApplicationKey' $json_file)
    read -r -p "$(tput setaf 6)Input a Datadog Application Key [$default]: $(tput sgr 0)" datadogApplicationKey </dev/tty
    datadogApplicationKey=${datadogApplicationKey:-$default}

    configStr=$( jq -c $(printf '.dataAdapterSecret.datadogAPIKey="%s"' $datadogAPIKey) <<<$configStr )
    configStr=$( jq -c $(printf '.dataAdapterSecret.datadogApplicationKey="%s"' $datadogApplicationKey) <<<$configStr )
    jq -r '.' <<< $configStr > $json_file
}

get_kafka_info()
{
    index=0
    next_set="y"
    while [[ "$next_set" = "y" ]]
    do
        echo -e "\nGetting the Kafka info... No.$((index+1))"
        default=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerDeploymentName' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer deployment name [$default]: $(tput sgr 0)" kafkaConsumerDeploymentName </dev/tty
        kafkaConsumerDeploymentName=${kafkaConsumerDeploymentName:-$default}

        default=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerDeploymentNamespace' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer deeployment namespace [$default]: $(tput sgr 0)" kafkaConsumerDeploymentNamespace </dev/tty
        kafkaConsumerDeploymentNamespace=${kafkaConsumerDeploymentNamespace:-$default}

        default=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerMinimumReplica' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer minimum replica number [$default]: $(tput sgr 0)" kafkaConsumerMinimumReplica </dev/tty
        kafkaConsumerMinimumReplica=${kafkaConsumerMinimumReplica:-$default}

        default=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerMaximumReplica' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer maximum replica number [$default]: $(tput sgr 0)" kafkaConsumerMaximumReplica </dev/tty
        kafkaConsumerMaximumReplica=${kafkaConsumerMaximumReplica:-$default}

        default=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerGroupName' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer group name [$default]: $(tput sgr 0)" kafkaConsumerGroupName </dev/tty
        kafkaConsumerGroupName=${kafkaConsumerGroupName:-$default}

        default=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaConsumerGroupNamespace' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer group namespace [$default]: $(tput sgr 0)" kafkaConsumerGroupNamespace </dev/tty
        kafkaConsumerGroupNamespace=${kafkaConsumerGroupNamespace:-$default}

        default=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaTopicName' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer topic name [$default]: $(tput sgr 0)" kafkaTopicName </dev/tty
        kafkaTopicName=${kafkaTopicName:-$default}

        default=$(jq -r '.dataAdapterConfigmap['$index'] | .kafkaTopicNamespace' $json_file | sed 's:null::g')
        read -r -p "$(tput setaf 6)Input Kafka consumer topic namespace [$default]: $(tput sgr 0)" kafkaTopicNamespace </dev/tty
        kafkaTopicNamespace=${kafkaTopicNamespace:-$default}

        if [ "$kafkaConsumerDeploymentName" = "" ] || [ "$kafkaConsumerDeploymentNamespace" = "" ] || [ "$kafkaConsumerMinimumReplica" = "" ] || [ "$kafkaConsumerMaximumReplica" = "" ] || [ "$kafkaConsumerGroupName" = "" ] || [ "$kafkaConsumerGroupNamespace" = "" ] || [ "$kafkaTopicName" = "" ] || [ "$kafkaTopicNamespace" = "" ]; then
            echo -e "\n$(tput setaf 1)Error! Kafka info can't be empty.$(tput sgr 0)"
            exit 7
        fi

        configStr=$( jq -c $(printf '.dataAdapterConfigmap[%s].kafkaConsumerDeploymentName="%s"' $index $kafkaConsumerDeploymentName) <<<$configStr )
        configStr=$( jq -c $(printf '.dataAdapterConfigmap[%s].kafkaConsumerDeploymentNamespace="%s"' $index $kafkaConsumerDeploymentNamespace) <<<$configStr )
        configStr=$( jq -c $(printf '.dataAdapterConfigmap[%s].kafkaConsumerMinimumReplica="%s"' $index $kafkaConsumerMinimumReplica) <<<$configStr )
        configStr=$( jq -c $(printf '.dataAdapterConfigmap[%s].kafkaConsumerMaximumReplica="%s"' $index $kafkaConsumerMaximumReplica) <<<$configStr )
        configStr=$( jq -c $(printf '.dataAdapterConfigmap[%s].kafkaConsumerGroupName="%s"' $index $kafkaConsumerGroupName) <<<$configStr )
        configStr=$( jq -c $(printf '.dataAdapterConfigmap[%s].kafkaConsumerGroupNamespace="%s"' $index $kafkaConsumerGroupNamespace) <<<$configStr )
        configStr=$( jq -c $(printf '.dataAdapterConfigmap[%s].kafkaTopicName="%s"' $index $kafkaTopicName) <<<$configStr )
        configStr=$( jq -c $(printf '.dataAdapterConfigmap[%s].kafkaTopicNamespace="%s"' $index $kafkaTopicNamespace) <<<$configStr )
        jq -r '.' <<< $configStr > $json_file
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
            configStr=$( jq -c $(printf 'del(.dataAdapterConfigmap[%s:])' $index) <<<$configStr )
            jq -r '.' <<< $configStr > $json_file
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

prepare_env()
{
    get_alamedaservice_full_version

    if [ ! -f "$json_file" ]; then
        if [ ! -f "$json_file_template" ]; then
            if ! curl -sL --fail https://raw.githubusercontent.com/containers-ai/federatorai-operator/${alameda_version}/deploy/${json_file_template} -O; then
                echo -e "\n$(tput setaf 1)Abort, download $json_file_template file failed!!!$(tput sgr 0)"
                exit 2
            fi
        fi
        cp $json_file_template $json_file
    fi

    configStr=$(jq -c '.' $json_file)
    which jq > /dev/null 2>&1
    if [ "$?" != "0" ];then
        if ! curl -sL --fail https://raw.githubusercontent.com/containers-ai/federatorai-operator/4.2-husky/deploy/jq -O; then
            echo -e "\n$(tput setaf 1)Abort, download jq binary failed!!!$(tput sgr 0)"
            exit 2
        fi
        chmod +x jq
        mv jq /usr/bin/
    fi
}

if [ "$#" -eq "0" ]; then
    show_usage
    exit 1
fi

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

file_folder="./config_result"
current_location=`pwd`
mkdir -p $file_folder

if [ "${kubeconfig}" = "" ]; then
    echo -e "\n$(tput setaf 1)Error! Need to use \"-k\" to specify openshift kubeconfig file.$(tput sgr 0)"
    show_usage
fi
export KUBECONFIG=${kubeconfig}
# Check if kubectl connect to server.
result="`echo ""|kubectl cluster-info 2>/dev/null`"
if [ "$?" != "0" ]; then
    echo -e "\n$(tput setaf 1)Error! Please login into OpenShift cluster first.$(tput sgr 0)"
    exit 1
fi
current_server="`echo $result|sed 's/.*at //'|awk '{print $1}'`"
echo "You are connecting to cluster: $current_server"

install_namespace="`kubectl get pods --all-namespaces |grep "alameda-ai-"|awk '{print $1}'|head -1`"
if [ "$install_namespace" = "" ];then
    echo -e "\n$(tput setaf 1)Error! Please install Federatorai before running this script.$(tput sgr 0)"
    exit 3
fi

json_file="adapter.json"
json_file_template="adapter.json.tmp"

prepare_env

get_datadog_key

get_kafka_info

patch_data_adapter_secret

patch_data_adapter_configmap

echo -e "$(tput setaf 6)\nSetup Federator.ai for datadog successfully$(tput sgr 0)"

exit 0