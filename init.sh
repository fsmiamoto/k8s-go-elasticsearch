#/bin/sh
#
# Run this script for creating the resources on kubernetes
#

kubectl create -f ./elastic/elasticsearch.yaml
kubectl create -f ./elastic/apm-server.yaml
kubectl create -f ./elastic/kibana.yaml

eval $(minikube docker-env)

docker build app/src/ -t k8s-golang-elastic-app

eval $(minikube docker-env -u)

kubectl create -f ./app/app.yaml
