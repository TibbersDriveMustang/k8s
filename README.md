
# Start a local kubernetes cluster
>minikube start

# Direct kubectl to use minikube(Optional, if not auto switched)
>kubectl config use-context minikube

# Start Redis Cluster:

>kubectl apply -f redis-sts.yaml

>kubectl apply -f redis-svc.yaml

# Setup Redis Cluster, Set first redis pod as master:
>kubectl exec -it redis-cluster-0 redis-cli cluster nodes

### **Issue, cluster nodes not auto connected**

# Verify Redis-cluster pods are running:
>kubectl get pods

# Verify Redis-cluster Service is running:
>kubectl get services

# Verify one of the redis pod is running
>kubectl exec -it redis-cluster-0 redis-cli ping

# Access one of the Redis Pod:
>kubectl exec -it redis-cluster-0 redis-cli

## Access Redis-cluster via minikube
>minikube service redis-cluster


# Clean Up
## Delete app

## Delete service
>kubectl delete service redis-cluster

## Delete configmaps

## Delete redis-cluster
>

## Delete minikube(Optional)
>minikube delete
