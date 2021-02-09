package cmd

import (
	"tubectl/api"

	"github.com/spf13/cobra"
)

func init() {
	//fmt.Println("in deployment")
	createCmd.AddCommand(CreateDeploymentCmd)
	deleteCmd.AddCommand(DeleteDeploymentCmd)
	getCmd.AddCommand(GetDeploymentCmd)
	updateCmd.AddCommand(UpdateDeploymentCmd)
	listCmd.AddCommand(ListDeploymentCmd)
}

var CreateDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "deployment is a sub-command of create command, it is used to ",
	Long:  "'create-deployment' command is for creating a deployment object using k8s API",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deployment called")
		api.CreateDeployment()
	},
}

var DeleteDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "deployment is a sub-command of create command, it is used to ",
	Long:  "'create-deployment' command is for creating a deployment object using k8s API",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deployment called")
		api.DeleteDeployment()
	},
}

var GetDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "deployment is a sub-command of create command, it is used to ",
	Long:  "'create-deployment' command is for creating a deployment object using k8s API",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deployment called")
		api.GetDeployment()
	},
}

var UpdateDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "deployment is a sub-command of create command, it is used to ",
	Long:  "'create-deployment' command is for creating a deployment object using k8s API",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deployment called")
		api.UpdateDeployment()
	},
}

var ListDeploymentCmd = &cobra.Command{
	Use:   "deployment",
	Short: "deployment is a sub-command of create command, it is used to ",
	Long:  "'create-deployment' command is for creating a deployment object using k8s API",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("deployment called")
		api.ListDeployment()
	},
}
