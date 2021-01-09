package test

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/organizations"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)
var terraformPlan = true


func TestOrganizationAccountCreation(t *testing.T) {

	// Set up expected values to be checked later
	expectedAccountName := fmt.Sprintf("a4l-dev-%s", random.UniqueId())
	expectedEmail := fmt.Sprintf("mnsanfilippo+dev+%s@gmail.com", random.UniqueId())

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{

		// The path to where our Terraform code is located
		TerraformDir: "./../../accounts/dev",

		//Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"account_name":  expectedAccountName,
			"account_email": expectedEmail,
		},

		////Environment variables to set when running Terraform
		//EnvVars: map[string]string{
		//	"AWS_DEFAULT_REGION": os,
		//},
	})

	//At the end of the test, run "terraform destroy" to clean up any resources that were created
	// In this case, the account is not deleted, to be deleted you need to complete first the account sign-up steps
	//defer terraform.Destroy(t, terraformOptions)

	if terraformPlan {
		// Also, you can only do terraform plan
		terraform.Init(t, terraformOptions)
		terraform.Plan(t, terraformOptions)

	} else {
		// This will run `terraform init` and `terraform apply`and fail the test if there any errors
		terraform.InitAndApply(t, terraformOptions)

		// Look up the Organization Account by id
		accountId := terraform.Output(t, terraformOptions, "account_id")
		account := describeAccountById(accountId)

		assert.Equal(t, expectedEmail, *account.Email)
		assert.Equal(t, expectedAccountName, *account.Name)

		// I added this only to see if the account was created in the AWS console before it was deleted
		log.Println("Now you can go and check if the account was created in the AWS Console.\n" +
			"Remember, to delete the account, you firs have to complete the sign-up process.")
	}
}

func describeAccountById(id string) *types.Account {

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	svc := organizations.NewFromConfig(cfg)

	req, err := svc.DescribeAccount(context.Background(), &organizations.DescribeAccountInput{AccountId: aws.String(id)})
	if err != nil {
		log.Println("Failed retrieving account by ID")
		log.Fatal(err)
	}

	return req.Account
}


