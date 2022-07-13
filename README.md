Kubernetes provider server-apply and patching trials

This has examples/trials of using the Pulumi Kubernetes provider with server-side apply,
and using *Patch objects.

### patch-configmap

This has an "original" config map, holding important configuration, in `original.yaml`. The Pulumi
program patches the config with another value.

**NB**:

 - trying to patch the same data key leads to a conflict

