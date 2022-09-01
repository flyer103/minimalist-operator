The project is designed to implement a K8s Operator with as few tools as possible and understand the K8s Operator core logic.

# How to write an Operator

1. Write CRD and register CR with kube-apiserver: [crd.yaml](./yaml/crd.yaml)
2. Write resource definitions through code: ![apis.png](./misc/apis/png)
3. Generate clients:
```
$ make build-resource
```
4. Write controller and add event handlers to informer.


# Usage

```shell
# Register CR.
$ kubectl apply -f yaml/crd.yaml

# Build Operator.
$ make build-operator

# Run operator outside of Cluster.
$ ./release/operator -kubeconfig ~/.kube/config
```
