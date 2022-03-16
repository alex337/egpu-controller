/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	egpuController "github.com/alex337/egpu-controller/pkg/apis/egpuController/v1alpha1"
	"flag"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"

	clientset "github.com/alex337/egpu-controller/pkg/generated/clientset/versioned"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}
	result, err := exampleClient.EgpucontrollerV1alpha1().EGPUs(apiv1.NamespaceDefault).Create(context.TODO(),newEGPU("test"), metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	klog.Info("Created EGPU:", result.GetObjectMeta().GetName())

}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func newEGPU(name string) *egpuController.EGPU {
	return &egpuController.EGPU{
		TypeMeta: metav1.TypeMeta{APIVersion: egpuController.SchemeGroupVersion.String()},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: metav1.NamespaceDefault,
		},
		Spec: egpuController.EGPUSpec{
			NodeName: name,
			GPU: []string{"1","2"},
			Resources: egpuController.EGPUResource{
				Capacity: egpuController.EGPUCapacity{
					QGPUMemory: "1",
					QGPUCore: "1",
				},
			},
		},
	}
}

func int32Ptr(i int32) *int32 { return &i }
