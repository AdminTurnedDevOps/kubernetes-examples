package main

import (
	"fmt"

	appsv1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/apps/v1"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		conf := config.New(ctx, "")

		appLabels := pulumi.StringMap{
			"app": pulumi.String(conf.Require("appName")),
		}
		deployment, err := appsv1.NewDeployment(ctx, conf.Require("deployment"), &appsv1.DeploymentArgs{
			Spec: appsv1.DeploymentSpecArgs{
				Selector: &metav1.LabelSelectorArgs{
					MatchLabels: appLabels,
				},
				Replicas: pulumi.Int(2),
				Template: &corev1.PodTemplateSpecArgs{
					Metadata: &metav1.ObjectMetaArgs{
						Labels: appLabels,
					},
					Spec: &corev1.PodSpecArgs{
						Containers: corev1.ContainerArray{
							corev1.ContainerArgs{
								Name:  pulumi.String(conf.Require("containerName")),
								Image: pulumi.String(conf.Require("imageName")),
							}},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		service, err := corev1.NewService(ctx, conf.Require("service"), &corev1.ServiceArgs{
			Spec: &corev1.ServiceSpecArgs{
				Ports: corev1.ServicePortArray{
					&corev1.ServicePortArgs{
						Port:     pulumi.Int(80),
						Protocol: pulumi.String("TCP"),
					},
				},
				Selector: pulumi.StringMap{
					"app": pulumi.String(conf.Require("appName")),
				},
			},
		})

		if err != nil {
			return err
		}

		ctx.Export("name", deployment.Metadata.Elem().Name())
		fmt.Println(service)
		// ctx.Export("name", service.Metadata.Elem().Name())

		return nil
	})
}
