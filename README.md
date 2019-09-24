# Start Redis Cluster:
>kubectl apply -f redis-sts.yaml

>kubectl apply -f redis-svc.yaml

# Verify Redis-cluster pods are running:
>kubectl get pods


# Verify Redis-cluster Service is running:
>kubectl get services

# Verify one of the redis pod is running:
>kubectl exec -it redis-cluster-0 redis-cli ping

# Access one of the Redis Pod:
>kubectl exec -it redis-cluster-0 redis-cli

## Access Redis-cluster via minikube
>minikube service redis-cluster
