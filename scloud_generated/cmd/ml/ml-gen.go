// Package ml -- generated by scloudgen
// !! DO NOT EDIT !! 
// 
package ml

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/scloud_generated/pkg/ml"
)


// createWorkflow -- 
var createWorkflowCmd = &cobra.Command{
	Use:   "create-workflow",
	Short: "",
	RunE:  impl.CreateWorkflow,
}

// createWorkflowBuild -- 
var createWorkflowBuildCmd = &cobra.Command{
	Use:   "create-workflow-build",
	Short: "",
	RunE:  impl.CreateWorkflowBuild,
}

// createWorkflowDeployment -- 
var createWorkflowDeploymentCmd = &cobra.Command{
	Use:   "create-workflow-deployment",
	Short: "",
	RunE:  impl.CreateWorkflowDeployment,
}

// createWorkflowInference -- 
var createWorkflowInferenceCmd = &cobra.Command{
	Use:   "create-workflow-inference",
	Short: "",
	RunE:  impl.CreateWorkflowInference,
}

// createWorkflowRun -- 
var createWorkflowRunCmd = &cobra.Command{
	Use:   "create-workflow-run",
	Short: "",
	RunE:  impl.CreateWorkflowRun,
}

// deleteWorkflow -- 
var deleteWorkflowCmd = &cobra.Command{
	Use:   "delete-workflow",
	Short: "",
	RunE:  impl.DeleteWorkflow,
}

// deleteWorkflowBuild -- 
var deleteWorkflowBuildCmd = &cobra.Command{
	Use:   "delete-workflow-build",
	Short: "",
	RunE:  impl.DeleteWorkflowBuild,
}

// deleteWorkflowDeployment -- 
var deleteWorkflowDeploymentCmd = &cobra.Command{
	Use:   "delete-workflow-deployment",
	Short: "",
	RunE:  impl.DeleteWorkflowDeployment,
}

// deleteWorkflowRun -- 
var deleteWorkflowRunCmd = &cobra.Command{
	Use:   "delete-workflow-run",
	Short: "",
	RunE:  impl.DeleteWorkflowRun,
}

// getWorkflow -- 
var getWorkflowCmd = &cobra.Command{
	Use:   "get-workflow",
	Short: "",
	RunE:  impl.GetWorkflow,
}

// getWorkflowBuild -- 
var getWorkflowBuildCmd = &cobra.Command{
	Use:   "get-workflow-build",
	Short: "",
	RunE:  impl.GetWorkflowBuild,
}

// getWorkflowBuildError -- 
var getWorkflowBuildErrorCmd = &cobra.Command{
	Use:   "get-workflow-build-error",
	Short: "",
	RunE:  impl.GetWorkflowBuildError,
}

// getWorkflowBuildLog -- 
var getWorkflowBuildLogCmd = &cobra.Command{
	Use:   "get-workflow-build-log",
	Short: "",
	RunE:  impl.GetWorkflowBuildLog,
}

// getWorkflowDeployment -- 
var getWorkflowDeploymentCmd = &cobra.Command{
	Use:   "get-workflow-deployment",
	Short: "",
	RunE:  impl.GetWorkflowDeployment,
}

// getWorkflowDeploymentError -- 
var getWorkflowDeploymentErrorCmd = &cobra.Command{
	Use:   "get-workflow-deployment-error",
	Short: "",
	RunE:  impl.GetWorkflowDeploymentError,
}

// getWorkflowDeploymentLog -- 
var getWorkflowDeploymentLogCmd = &cobra.Command{
	Use:   "get-workflow-deployment-log",
	Short: "",
	RunE:  impl.GetWorkflowDeploymentLog,
}

// getWorkflowRun -- 
var getWorkflowRunCmd = &cobra.Command{
	Use:   "get-workflow-run",
	Short: "",
	RunE:  impl.GetWorkflowRun,
}

// getWorkflowRunError -- 
var getWorkflowRunErrorCmd = &cobra.Command{
	Use:   "get-workflow-run-error",
	Short: "",
	RunE:  impl.GetWorkflowRunError,
}

// getWorkflowRunLog -- 
var getWorkflowRunLogCmd = &cobra.Command{
	Use:   "get-workflow-run-log",
	Short: "",
	RunE:  impl.GetWorkflowRunLog,
}

// listWorkflowBuilds -- 
var listWorkflowBuildsCmd = &cobra.Command{
	Use:   "list-workflow-builds",
	Short: "",
	RunE:  impl.ListWorkflowBuilds,
}

// listWorkflowDeployments -- 
var listWorkflowDeploymentsCmd = &cobra.Command{
	Use:   "list-workflow-deployments",
	Short: "",
	RunE:  impl.ListWorkflowDeployments,
}

// listWorkflowRuns -- 
var listWorkflowRunsCmd = &cobra.Command{
	Use:   "list-workflow-runs",
	Short: "",
	RunE:  impl.ListWorkflowRuns,
}

// listWorkflows -- 
var listWorkflowsCmd = &cobra.Command{
	Use:   "list-workflows",
	Short: "",
	RunE:  impl.ListWorkflows,
}


func init() {

    mlCmd.AddCommand(createWorkflowCmd)
    mlCmd.AddCommand(createWorkflowBuildCmd)
    mlCmd.AddCommand(createWorkflowDeploymentCmd)
    mlCmd.AddCommand(createWorkflowInferenceCmd)
    mlCmd.AddCommand(createWorkflowRunCmd)
    mlCmd.AddCommand(deleteWorkflowCmd)
    mlCmd.AddCommand(deleteWorkflowBuildCmd)
    mlCmd.AddCommand(deleteWorkflowDeploymentCmd)
    mlCmd.AddCommand(deleteWorkflowRunCmd)
    mlCmd.AddCommand(getWorkflowCmd)
    mlCmd.AddCommand(getWorkflowBuildCmd)
    mlCmd.AddCommand(getWorkflowBuildErrorCmd)
    mlCmd.AddCommand(getWorkflowBuildLogCmd)
    mlCmd.AddCommand(getWorkflowDeploymentCmd)
    mlCmd.AddCommand(getWorkflowDeploymentErrorCmd)
    mlCmd.AddCommand(getWorkflowDeploymentLogCmd)
    mlCmd.AddCommand(getWorkflowRunCmd)
    mlCmd.AddCommand(getWorkflowRunErrorCmd)
    mlCmd.AddCommand(getWorkflowRunLogCmd)
    mlCmd.AddCommand(listWorkflowBuildsCmd)
    mlCmd.AddCommand(listWorkflowDeploymentsCmd)
    mlCmd.AddCommand(listWorkflowRunsCmd)
    mlCmd.AddCommand(listWorkflowsCmd)
    

	// subTest1Cmd.Flags().StringP("id", "i", "", "resource identifier")
	// subTest2Cmd.Flags().StringP("id", "i", "", "resource identifier")
}