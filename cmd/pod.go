package cmd

import (
	"tubectl/api"

	"github.com/spf13/cobra"
)

var podName string
var imageName string

func init() {
	createCmd.AddCommand(createPodCmd)
	getCmd.AddCommand(getPodCmd)
	deleteCmd.AddCommand(deletePodCmd)
	updateCmd.AddCommand(updatePodCmd)

	updateCmd.PersistentFlags().StringVarP(&podName, "podname", "p", "", "this is for getting a podName for updating that pod")
	updateCmd.PersistentFlags().StringVarP(&imageName, "imagename", "n", "", "this is for getting a updated image of a pod")
}

// createPodCmd represents the pod command
var createPodCmd = &cobra.Command{
	Use:   "pod",
	Short: "this command is for creating a pod in default namespace",
	Long:  "this command is for creating a pod in default namespace",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("pod called")
		api.CreatePod()
	},
}

// getPodCmd represents the pods command
var getPodCmd = &cobra.Command{
	Use:   "pods",
	Short: "this command is for getting all the pods in default namespace",
	Long:  "this command is for getting all the pods in default namespace",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("pod called")
		api.GetPods()
	},
}

// deletePodCmd represents the pod command
var deletePodCmd = &cobra.Command{
	Use:   "pod",
	Short: "this command is for deleting a pod in default namespace",
	Long:  "this command is for deleting a pod in default namespace",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("pod called")
		api.DeletePods(args)
	},
}

// updatePodCmd represents the pod command
var updatePodCmd = &cobra.Command{
	Use:   "pod",
	Short: "this command is for updating a pod in default namespace",
	Long:  "this command is for updating a pod in default namespace",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("pod called")
		api.UpdatePod(podName, imageName)
	},
}
