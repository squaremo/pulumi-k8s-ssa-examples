package main

import (
	k8s "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		provider, err := k8s.NewProvider(ctx, "ssa", &k8s.ProviderArgs{
			Kubeconfig:            pulumi.StringPtr("/Users/mikeb/.kube/config"),
			EnableServerSideApply: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		_, err = corev1.NewConfigMap(ctx, "bar", &corev1.ConfigMapArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String("bar"),
				Annotations: pulumi.StringMap{
					"pulumi.com/patchForce": pulumi.String("true"),
				},
			},
			Data: pulumi.StringMap{
				"bar": pulumi.String("bing"),
			},
		}, pulumi.Provider(provider))
		if err != nil {
			return err
		}

		return nil
	})
}
