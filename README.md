# operator-example

An example Kubernetes operator that creates pods
The operator created with the `operator-sdk` tool

## Steps for creation
```
operator-sdk init --domain nati.com --repo operator-example
operator-sdk create api --group=example --version=v1 --kind=DemoPod
```

Set the spec in the api/v1 types.go file and generate the CRD using `make` command:
```
make generate
make manifests
```

## Install the CRD
Install the CRD using `make` command:
```
make install
```

A sample file created in config/samples, you can use it to run your operator:
```
kubectl create -f config/samples/example_v1_demopod.yaml
```

## Check the CRD installed
You can get the installed CRDs using `kubectl` command:
```
kubectl get crd
kubectl get demopods.example.nati.com
```

## Running the operator
You can run the operator using `make` command:
```
make run
```
