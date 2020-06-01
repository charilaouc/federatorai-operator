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

ask_datadog_credential()
{
    read -r -p "$(tput setaf 6)Please input Datadog API Key: $(tput sgr 0): " datadog_api_key </dev/tty
    datadog_api_key=`echo -n "$datadog_api_key" | base64`
}

patch_data_adapter_secret()
{
    ask_datadog_credential
}

get_datadog_key()
{

}

get_kafka_info()
{

}

prepare_env()
{
    if [ ! -f "$json_file" ]; then
        if [ ! -f "$json_file_template" ]; then
            if ! curl -sL --fail https://raw.githubusercontent.com/containers-ai/federatorai-operator/master/deploy/${json_file_template} -O; then
                echo -e "\n$(tput setaf 1)Abort, download $json_file_template file failed!!!$(tput sgr 0)"
                exit 2
            fi
        fi
        cp $json_file_template $json_file
    fi

    configstr=$(jq -c '.' $json_file)
    which jq > /dev/null 2>&1
    if [ "$?" != "0" ];then
        if ! curl -sL --fail https://raw.githubusercontent.com/containers-ai/federatorai-operator/master/deploy/jq -O; then
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
json_file_template="adapter.json.template"

prepare_env

get_datadog_key

get_kafka_info

patch_data_adapter_secret
#### get datadog api key
#### get datadog application key

#repeat
#### kafka_consumer_deployment_name
#### kafka_consumer_deployment_namespace
#### min_replicas
#### max_replicas
#### kafka_consumer_group_name
#### kafka_consumer_group_namespace
#### kafka_topic_name
#### kafka_topic_namespace

# skip for now
#### cluster name
#### 
#### 
