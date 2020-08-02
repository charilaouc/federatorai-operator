#!/usr/bin/env bash

leave_prog()
{
    echo -e "\n$(tput setaf 5)Downloaded script files are located under $file_folder $(tput sgr 0)"
    cd $current_location > /dev/null
}

which curl > /dev/null 2>&1
if [ "$?" != "0" ];then
    echo -e "\n$(tput setaf 1)Abort, \"curl\" command is needed for this tool.$(tput sgr 0)"
    exit
fi

while [ "$pass" != "y" ]   
do
    read -r -p "$(tput setaf 2)Please input Federator.ai version tag (e.g., v4.2.755): $(tput sgr 0) " tag_number </dev/tty
    if [[ $tag_number =~ ^[v][[:digit:]]+.[[:digit:]]+.[[:digit:]]+$ ]]; then
        pass="y"
    fi
    if [ "$pass" != "y" ]; then
        echo -e "\n$(tput setaf 1)Error, the tag_number should follow the correct format (e.g., v4.2.755).$(tput sgr 0)"
    fi
done

echo $tag_number|grep -q "v4.3" && version="4.3"
echo $tag_number|grep -q "v4.2" && version="4.2"
echo $tag_number|grep -q "v4.5" && version="4.5"

file_folder="/tmp/federatorai-scripts/${tag_number}"
rm -rf $file_folder
mkdir -p $file_folder
current_location=`pwd`
cd $file_folder

filearray=("install.sh" "email-notifier-setup.sh" "node-label-assignor.sh" "planning-util.sh" "preloader-util.sh" "prepare-private-repository.sh" "uninstall.sh")

if [ "$version" = "4.3" ]; then
    filearray=("${filearray[@]}" "federatorai-setup-for-datadog.sh")
fi

cr_files=( "alamedascaler.yaml" "alamedadetection.yaml" "alamedanotificationchannel.yaml" "alamedanotificationtopic.yaml" )

echo -e "\n$(tput setaf 6)Downloading scripts ...$(tput sgr 0)"

for file_name in "${filearray[@]}"
do
    if ! curl -sL --fail https://raw.githubusercontent.com/containers-ai/federatorai-operator/${tag_number}/deploy/${file_name} -O; then
        echo -e "\n$(tput setaf 1)Abort, download file $file_name failed!!!$(tput sgr 0)"
        echo "Please check tag name and network"
        exit 1
    fi
done

echo "Done"

while [ "$enable_private_repo" != "y" ] && [ "$enable_private_repo" != "n" ]    
do
    default="n"
    read -r -p "$(tput setaf 2)Do you want to use private repository URL? [default: $default]: $(tput sgr 0)" enable_private_repo </dev/tty
    enable_private_repo=${enable_private_repo:-$default}
    enable_private_repo=$(echo "$enable_private_repo" | tr '[:upper:]' '[:lower:]')
done

if [ "$enable_private_repo" = "y" ]; then
    read -r -p "$(tput setaf 2)Please input private repository URL (e.g., repo.prophetstor.com/federatorai): $(tput sgr 0) " repo_url </dev/tty
    repo_url=$(echo "$repo_url" | tr '[:upper:]' '[:lower:]')
    bash prepare-private-repository.sh $tag_number $repo_url
    # For install.sh
    export RELATED_IMAGE_URL_PREFIX=$repo_url
fi

echo -e "\n$(tput setaf 6)Executing install.sh ...$(tput sgr 0)"
bash install.sh -t $tag_number

leave_prog



