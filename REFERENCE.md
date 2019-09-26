# Expose the service
>kubectl expose deployment hit-counter-app --port=8080 --type=NodePort

# Setup Simple test app
>kubectl apply -f app-hit-counter.yaml

## Access Redis-cluster via minikube(For Reference)
>minikube service redis-cluster

# Get Node IP
>minikube ip
