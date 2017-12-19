# Client command

```
while true; do for i in {1..10}; do (curl http://localhost:8085/ &); done; sleep 2; done;
```

# deployment

```
unset KUBECONFIG
kubectl apply -f artifact/hello.svc.yaml
kubectl apply -f <(istioctl kube-inject -f artifact/hello.deployment.yaml)
kubectl apply -f <(istioctl kube-inject -f artifact/inject.deployment.yaml)

```

#screening connections
```
podname=$(k get pod -l run=inject -ojsonpath={.items[0].metadata.name})
helloip=$(k get pod -l run=hello -ojsonpath={.items[0].status.podIP})
```
