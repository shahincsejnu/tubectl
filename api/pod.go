package api

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/util/retry"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// CreatePod creates a pod
func CreatePod() {
	var clientset kubernetes.Interface
	clientset = CreateClient()

	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

	pod := &apiv1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "pod-demo",
			Labels: map[string]string{
				"app": "demo-pod",
			},
		},
		Spec: apiv1.PodSpec{
			Containers: []apiv1.Container{
				{
					Name:  "web-container",
					Image: "nginx:1.14",
					Ports: []apiv1.ContainerPort{
						{
							Name:          "http-oka",
							Protocol:      apiv1.ProtocolTCP,
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}

	// Create pod
	fmt.Println("Creating pod...")
	result, err := podClient.Create(context.TODO(), pod, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}
	fmt.Printf("Created pod %q.\n", result.GetObjectMeta().GetName())
}

// GetPods for getting all the pods in default namespace
func GetPods() {
	var clientset kubernetes.Interface
	clientset = CreateClient()

	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

	//fmt.Printf("Getting all the pods in namespace %q:\n", apiv1.NamespaceDefault)

	list, err := podClient.List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Name\t\t\tREADY\t\tSTATUS\t\tRESTARTS\tAGE")
	for _, d := range list.Items {
		end := time.Now()
		start := d.CreationTimestamp.Time
		diff := int(end.Sub(start).Seconds())
		hour := diff / 3600
		diff %= 3600
		minute := diff / 60
		diff %= 60
		second := diff
		//hour := int(diff.Hours())
		//minute := int(diff.Minutes())
		//second := int(diff.Seconds())

		//diff = diff.Truncate(60)
		fmt.Printf("%s\t\t%d/%d\t\t%v\t\t%d\t\t%dh:%dm:%ds\n", d.Name, 1, 1, d.Status.Phase, d.Status.ContainerStatuses[0].RestartCount, hour, minute, second)
	}
}

//DeletePods for deleting pods
func DeletePods(args []string) {
	var clientset kubernetes.Interface
	clientset = CreateClient()

	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

	deletePolicy := metav1.DeletePropagationForeground

	for _, podName := range args {
		if err := podClient.Delete(context.TODO(), podName, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
		fmt.Printf("Deleted pod : %s\n", podName)
	}
}

//UpdatePod for updating a pod
func UpdatePod(podName string, imageName string) {
	var clientset kubernetes.Interface
	clientset = CreateClient()

	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := podClient.Get(context.TODO(), podName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of pod: %v", getErr))
		}

		result.Spec.Containers[0].Image = imageName
		_, updateErr := podClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Println("Updated the pod : ", podName)
}
