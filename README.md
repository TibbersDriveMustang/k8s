# First
>go get k8s.io/client-go@kubernetes-1.15.3

# Then
>go build -o app .

    Which downloads the rest dependencies    

# Install Docker for Mac
>https://docs.docker.com/docker-for-mac/install/

    installed v19.03

# Install Rancher:
>sudo docker run -d --restart=unless-stopped -p 80:80 -p 443:443 rancher/rancher  
>Go to htts://localhost to start a rancher Cluster 
>Link the running kubernetes cluster(raised by minikube) to rancher via import  

# Start Redis Cluster:
>kubectl apply -f redis-sts.yaml

>kubectl apply -f redis-svc.yaml

# Verify Redis-cluster pods:
>kubectl get pods

# Inspect single Pod(shows no thing):
>kubectl describe pods redis-cluster-0 | grep pvc

# Rancher status: pending
## Check cattle-node-agent status
>kubectl -n cattle-system get pods -l app=cattle-agent -o wide
#### log
>kubectl -n cattle-system logs -l app=cattle-agent

## Check kubernetes config
> vi ~/.kube/config

# Check redis-cluster serviceName
>kubectl get svc my-nginx
>kubectl describe svc redis-cluster
>kubectl exec redis-cluster-0 -- printenv | grep SERVICE

# Exposing Redis-cluster
>kubectl expose deployment hello --port=8080 --type=NodePort
## deployment not found
>kubectl expose deployment redis --type=LoadBalancer --name=redis


## Get Pod IP
>kubectl get pods -o yaml | grep -i podip

## Get minikue services
>minikube service redis-cluster

## Exec on Redis node
>kubectl exec -it redis-cluster-0 redis-cli
