package api

import (
	"context"
	"fmt"
	"time"

	"k8s.io/client-go/util/retry"

	"k8s.io/client-go/kubernetes"

	appsv1 "k8s.io/api/apps/v1"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateDeployment creates a deployment
func CreateDeployment() {
	var clientset kubernetes.Interface
	clientset = CreateClient()

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: "demo-deployment",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(2),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "demo",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "demo",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "nginx:1.12",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})

	if err != nil {
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
}

// GetDeployment for getting all the deployments in default namespace
func GetDeployment() {
	var clientset kubernetes.Interface
	clientset = CreateClient()

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	fmt.Printf("Getting all the deployments in namespace %q:\n", apiv1.NamespaceDefault)

	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Name\t\t\tREADY\t\tUP-TO-DATE\tAVAILABLE\tAGE")
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
		fmt.Printf("%s\t\t%d/%d\t\t%d\t\t%d\t\t%dh:%dm:%ds\n", d.Name, d.Status.ReadyReplicas, *d.Spec.Replicas, d.Status.UpdatedReplicas, d.Status.AvailableReplicas, hour, minute, second)
	}
}

//DeleteDeployment for deleteing a deployment
func DeleteDeployment(args []string) {
	var clientset kubernetes.Interface
	clientset = CreateClient()

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	deletePolicy := metav1.DeletePropagationForeground

	for _, deployName := range args {
		if err := deploymentsClient.Delete(context.TODO(), deployName, metav1.DeleteOptions{
			PropagationPolicy: &deletePolicy,
		}); err != nil {
			panic(err)
		}
		fmt.Printf("Deleted deployment : %s\n", deployName)
	}
}

//UpdateDeployment for updating a deployment
func UpdateDeployment(deployName string, replica int32, image string) {
	var clientset kubernetes.Interface
	clientset = CreateClient()

	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)

	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(context.TODO(), deployName, metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}

		result.Spec.Replicas = int32Ptr(replica)
		result.Spec.Template.Spec.Containers[0].Image = image
		_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	fmt.Println("Updated deployment...")
}
