Kubernetes provider server-apply and patching trials

This has examples/trials of using the Pulumi Kubernetes provider with server-side apply,
and using *Patch objects.

### patch-configmap

This has an "original" config map, holding important configuration, in `original.yaml`. The Pulumi
program patches the config with another value.

**NB**:

 - trying to patch the same data key leads to a conflict
   - you can use the annotation on the patch `pulumi.com/patchForce: true` to make it overwrite the other field
 - destroying the stack removes the extra field
   - if you forced an overwrite, the data key is removed

### namespace-meta

This patches the default namespace with some additional labels. Again, about as simple as you can get.

### ensure-configmap

This uses server-side apply to "upsert" a ConfigMap -- that is, to either create or overwrite a
particular ConfigMap, depending on whether it already exists or not.

Either create the original with `kubectl apply -f original.yaml` then run `pulumi up`, or skip the
first step.
