#!/usr/bin/env bash


#######################################
# This is Just a Reference shell Script
#######################################

# For Mac Install Command Line Tools in case you dont` have
xcode-select --install

#set -e

#sudo su

# TODO: Set to URL of git repo.
PROJECT_GIT_URL='https://github.com/TibbersDriveMustang/k8s.git'

#
sysctl -a | grep -E --color 'machdep.cpu.features|VMX' | grep 'VMX' &> /dev/null

if [ $? == 0 ]
then
    echo "virtualization is supported"
else
    echo "virtualization is not supported"
    exit 0
fi

# Download minikube
curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-darwin-amd64 \
  && chmod +x minikube

echo "Adding the Minikube to executable: "
echo "For Mac User.Please make sure System Integrity Protection is disabled"
echo "hint: csrutil disable"
sudo mv minikube /usr/local/bin


# If you have previously installed minikube
# And 'minikube start' returns error
# Do: 'minikube delete'

minikube start

echo "Make sure virtualBox is installed"

#Done! kubectl is now configured to use "minikube"
#For best results, install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/

#Install kubectl
    #Mac OS
    #Download lastest version
    curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl
    #Make executable
    chmod +x ./kubectl
    #Copy to path
    sudo cp ./kubectl /usr/local/bin/kubectl
    #Check version
    kubectl version

# For Mac, Install homebrew in case you don`t have
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"

# For Mac, Install wget in case you don`t have
brew install wget

echo "Installing Redis:"

wget http://download.redis.io/redis-stable.tar.gz
tar xvzf redis-stable.tar.gz
cd redis-stable
make



echo "DONE! :)"
