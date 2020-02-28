// Package appregistry -- generated by scloudgen
// !! DO NOT EDIT !!
//
package appregistry

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/flags"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/jsonx"
	model "github.com/splunk/splunk-cloud-sdk-go/services/appregistry"
)

// CreateAppNativeApp Creates an app.
func CreateAppNativeApp(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appPrincipalPermissions []string
	err = flags.ParseFlag(cmd.Flags(), "app-principal-permissions", &appPrincipalPermissions)
	if err != nil {
		return fmt.Errorf(`error parsing "app-principal-permissions": ` + err.Error())
	}
	var descriptionDefault string
	description := &descriptionDefault
	err = flags.ParseFlag(cmd.Flags(), "description", &description)
	if err != nil {
		return fmt.Errorf(`error parsing "description": ` + err.Error())
	}
	var kind model.AppResourceKind
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var loginUrlDefault string
	loginUrl := &loginUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "login-url", &loginUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "login-url": ` + err.Error())
	}
	var logoUrlDefault string
	logoUrl := &logoUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "logo-url", &logoUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "logo-url": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var redirectUrls []string
	err = flags.ParseFlag(cmd.Flags(), "redirect-urls", &redirectUrls)
	if err != nil {
		return fmt.Errorf(`error parsing "redirect-urls": ` + err.Error())
	}
	var setupUrlDefault string
	setupUrl := &setupUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "setup-url", &setupUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "setup-url": ` + err.Error())
	}
	var title string
	err = flags.ParseFlag(cmd.Flags(), "title", &title)
	if err != nil {
		return fmt.Errorf(`error parsing "title": ` + err.Error())
	}
	var userPermissionsFilter []string
	err = flags.ParseFlag(cmd.Flags(), "user-permissions-filter", &userPermissionsFilter)
	if err != nil {
		return fmt.Errorf(`error parsing "user-permissions-filter": ` + err.Error())
	}
	var webhookUrlDefault string
	webhookUrl := &webhookUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "webhook-url", &webhookUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-url": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.NativeAppPost{

		AppPrincipalPermissions: appPrincipalPermissions,
		Description:             description,
		Kind:                    kind,
		LoginUrl:                loginUrl,
		LogoUrl:                 logoUrl,
		Name:                    name,
		RedirectUrls:            redirectUrls,
		SetupUrl:                setupUrl,
		Title:                   title,
		UserPermissionsFilter:   userPermissionsFilter,
		WebhookUrl:              webhookUrl,
	}

	resp, err := client.AppRegistryService.CreateApp(model.MakeCreateAppRequestFromNativeAppPost(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreateAppServiceApp Creates an app.
func CreateAppServiceApp(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appPrincipalPermissions []string
	err = flags.ParseFlag(cmd.Flags(), "app-principal-permissions", &appPrincipalPermissions)
	if err != nil {
		return fmt.Errorf(`error parsing "app-principal-permissions": ` + err.Error())
	}
	var descriptionDefault string
	description := &descriptionDefault
	err = flags.ParseFlag(cmd.Flags(), "description", &description)
	if err != nil {
		return fmt.Errorf(`error parsing "description": ` + err.Error())
	}
	var kind model.AppResourceKind
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var loginUrlDefault string
	loginUrl := &loginUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "login-url", &loginUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "login-url": ` + err.Error())
	}
	var logoUrlDefault string
	logoUrl := &logoUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "logo-url", &logoUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "logo-url": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var redirectUrls []string
	err = flags.ParseFlag(cmd.Flags(), "redirect-urls", &redirectUrls)
	if err != nil {
		return fmt.Errorf(`error parsing "redirect-urls": ` + err.Error())
	}
	var setupUrlDefault string
	setupUrl := &setupUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "setup-url", &setupUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "setup-url": ` + err.Error())
	}
	var title string
	err = flags.ParseFlag(cmd.Flags(), "title", &title)
	if err != nil {
		return fmt.Errorf(`error parsing "title": ` + err.Error())
	}
	var userPermissionsFilter []string
	err = flags.ParseFlag(cmd.Flags(), "user-permissions-filter", &userPermissionsFilter)
	if err != nil {
		return fmt.Errorf(`error parsing "user-permissions-filter": ` + err.Error())
	}
	var webhookUrlDefault string
	webhookUrl := &webhookUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "webhook-url", &webhookUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-url": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.ServiceAppPost{

		AppPrincipalPermissions: appPrincipalPermissions,
		Description:             description,
		Kind:                    kind,
		LoginUrl:                loginUrl,
		LogoUrl:                 logoUrl,
		Name:                    name,
		RedirectUrls:            redirectUrls,
		SetupUrl:                setupUrl,
		Title:                   title,
		UserPermissionsFilter:   userPermissionsFilter,
		WebhookUrl:              webhookUrl,
	}

	resp, err := client.AppRegistryService.CreateApp(model.MakeCreateAppRequestFromServiceAppPost(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreateAppWebApp Creates an app.
func CreateAppWebApp(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appPrincipalPermissions []string
	err = flags.ParseFlag(cmd.Flags(), "app-principal-permissions", &appPrincipalPermissions)
	if err != nil {
		return fmt.Errorf(`error parsing "app-principal-permissions": ` + err.Error())
	}
	var descriptionDefault string
	description := &descriptionDefault
	err = flags.ParseFlag(cmd.Flags(), "description", &description)
	if err != nil {
		return fmt.Errorf(`error parsing "description": ` + err.Error())
	}
	var kind model.AppResourceKind
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	var loginUrlDefault string
	loginUrl := &loginUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "login-url", &loginUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "login-url": ` + err.Error())
	}
	var logoUrlDefault string
	logoUrl := &logoUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "logo-url", &logoUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "logo-url": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var redirectUrls []string
	err = flags.ParseFlag(cmd.Flags(), "redirect-urls", &redirectUrls)
	if err != nil {
		return fmt.Errorf(`error parsing "redirect-urls": ` + err.Error())
	}
	var setupUrlDefault string
	setupUrl := &setupUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "setup-url", &setupUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "setup-url": ` + err.Error())
	}
	var title string
	err = flags.ParseFlag(cmd.Flags(), "title", &title)
	if err != nil {
		return fmt.Errorf(`error parsing "title": ` + err.Error())
	}
	var userPermissionsFilter []string
	err = flags.ParseFlag(cmd.Flags(), "user-permissions-filter", &userPermissionsFilter)
	if err != nil {
		return fmt.Errorf(`error parsing "user-permissions-filter": ` + err.Error())
	}
	var webhookUrlDefault string
	webhookUrl := &webhookUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "webhook-url", &webhookUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-url": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.WebAppPost{

		AppPrincipalPermissions: appPrincipalPermissions,
		Description:             description,
		Kind:                    kind,
		LoginUrl:                loginUrl,
		LogoUrl:                 logoUrl,
		Name:                    name,
		RedirectUrls:            redirectUrls,
		SetupUrl:                setupUrl,
		Title:                   title,
		UserPermissionsFilter:   userPermissionsFilter,
		WebhookUrl:              webhookUrl,
	}

	resp, err := client.AppRegistryService.CreateApp(model.MakeCreateAppRequestFromWebAppPost(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// CreateSubscription Creates a subscription.
func CreateSubscription(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appName string
	err = flags.ParseFlag(cmd.Flags(), "app-name", &appName)
	if err != nil {
		return fmt.Errorf(`error parsing "app-name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.AppName{

		AppName: appName,
	}

	err = client.AppRegistryService.CreateSubscription(generated_request_body)
	if err != nil {
		return err
	}

	return nil
}

// DeleteApp Removes an app.
func DeleteApp(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appName string
	err = flags.ParseFlag(cmd.Flags(), "app-name", &appName)
	if err != nil {
		return fmt.Errorf(`error parsing "app-name": ` + err.Error())
	}

	err = client.AppRegistryService.DeleteApp(appName)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSubscription Removes a subscription.
func DeleteSubscription(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appName string
	err = flags.ParseFlag(cmd.Flags(), "app-name", &appName)
	if err != nil {
		return fmt.Errorf(`error parsing "app-name": ` + err.Error())
	}

	err = client.AppRegistryService.DeleteSubscription(appName)
	if err != nil {
		return err
	}

	return nil
}

// GetApp Returns the metadata of an app.
func GetApp(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appName string
	err = flags.ParseFlag(cmd.Flags(), "app-name", &appName)
	if err != nil {
		return fmt.Errorf(`error parsing "app-name": ` + err.Error())
	}

	resp, err := client.AppRegistryService.GetApp(appName)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetKeys Returns a list of the public keys used for verifying signed webhook requests.
func GetKeys(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClientSystemTenant()

	if err != nil {
		return err
	}

	resp, err := client.AppRegistryService.GetKeys()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetSubscription Returns or validates a subscription.
func GetSubscription(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appName string
	err = flags.ParseFlag(cmd.Flags(), "app-name", &appName)
	if err != nil {
		return fmt.Errorf(`error parsing "app-name": ` + err.Error())
	}

	resp, err := client.AppRegistryService.GetSubscription(appName)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListAppSubscriptions Returns the collection of subscriptions to an app.
func ListAppSubscriptions(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appName string
	err = flags.ParseFlag(cmd.Flags(), "app-name", &appName)
	if err != nil {
		return fmt.Errorf(`error parsing "app-name": ` + err.Error())
	}

	resp, err := client.AppRegistryService.ListAppSubscriptions(appName)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListApps Returns a list of apps.
func ListApps(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	resp, err := client.AppRegistryService.ListApps()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListSubscriptions Returns the tenant subscriptions.
func ListSubscriptions(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var kindDefault model.AppResourceKind
	kind := &kindDefault
	err = flags.ParseFlag(cmd.Flags(), "kind", &kind)
	if err != nil {
		return fmt.Errorf(`error parsing "kind": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListSubscriptionsQueryParams{}
	generated_query.Kind = kind

	resp, err := client.AppRegistryService.ListSubscriptions(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// RotateSecret Rotates the client secret for an app.
func RotateSecret(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appName string
	err = flags.ParseFlag(cmd.Flags(), "app-name", &appName)
	if err != nil {
		return fmt.Errorf(`error parsing "app-name": ` + err.Error())
	}

	resp, err := client.AppRegistryService.RotateSecret(appName)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UpdateAppCommonAppPut Updates an app.
func UpdateAppCommonAppPut(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appName string
	err = flags.ParseFlag(cmd.Flags(), "app-name", &appName)
	if err != nil {
		return fmt.Errorf(`error parsing "app-name": ` + err.Error())
	}
	var appPrincipalPermissions []string
	err = flags.ParseFlag(cmd.Flags(), "app-principal-permissions", &appPrincipalPermissions)
	if err != nil {
		return fmt.Errorf(`error parsing "app-principal-permissions": ` + err.Error())
	}
	var descriptionDefault string
	description := &descriptionDefault
	err = flags.ParseFlag(cmd.Flags(), "description", &description)
	if err != nil {
		return fmt.Errorf(`error parsing "description": ` + err.Error())
	}
	var loginUrlDefault string
	loginUrl := &loginUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "login-url", &loginUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "login-url": ` + err.Error())
	}
	var logoUrlDefault string
	logoUrl := &logoUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "logo-url", &logoUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "logo-url": ` + err.Error())
	}
	var redirectUrls []string
	err = flags.ParseFlag(cmd.Flags(), "redirect-urls", &redirectUrls)
	if err != nil {
		return fmt.Errorf(`error parsing "redirect-urls": ` + err.Error())
	}
	var setupUrlDefault string
	setupUrl := &setupUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "setup-url", &setupUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "setup-url": ` + err.Error())
	}
	var title string
	err = flags.ParseFlag(cmd.Flags(), "title", &title)
	if err != nil {
		return fmt.Errorf(`error parsing "title": ` + err.Error())
	}
	var userPermissionsFilter []string
	err = flags.ParseFlag(cmd.Flags(), "user-permissions-filter", &userPermissionsFilter)
	if err != nil {
		return fmt.Errorf(`error parsing "user-permissions-filter": ` + err.Error())
	}
	var webhookUrlDefault string
	webhookUrl := &webhookUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "webhook-url", &webhookUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-url": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.CommonAppPut{

		AppPrincipalPermissions: appPrincipalPermissions,
		Description:             description,
		LoginUrl:                loginUrl,
		LogoUrl:                 logoUrl,
		RedirectUrls:            redirectUrls,
		SetupUrl:                setupUrl,
		Title:                   title,
		UserPermissionsFilter:   userPermissionsFilter,
		WebhookUrl:              webhookUrl,
	}

	resp, err := client.AppRegistryService.UpdateApp(appName, model.MakeUpdateAppRequestFromCommonAppPut(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UpdateAppWebAppPut Updates an app.
func UpdateAppWebAppPut(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var appName string
	err = flags.ParseFlag(cmd.Flags(), "app-name", &appName)
	if err != nil {
		return fmt.Errorf(`error parsing "app-name": ` + err.Error())
	}
	var appPrincipalPermissions []string
	err = flags.ParseFlag(cmd.Flags(), "app-principal-permissions", &appPrincipalPermissions)
	if err != nil {
		return fmt.Errorf(`error parsing "app-principal-permissions": ` + err.Error())
	}
	var descriptionDefault string
	description := &descriptionDefault
	err = flags.ParseFlag(cmd.Flags(), "description", &description)
	if err != nil {
		return fmt.Errorf(`error parsing "description": ` + err.Error())
	}
	var loginUrlDefault string
	loginUrl := &loginUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "login-url", &loginUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "login-url": ` + err.Error())
	}
	var logoUrlDefault string
	logoUrl := &logoUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "logo-url", &logoUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "logo-url": ` + err.Error())
	}
	var redirectUrls []string
	err = flags.ParseFlag(cmd.Flags(), "redirect-urls", &redirectUrls)
	if err != nil {
		return fmt.Errorf(`error parsing "redirect-urls": ` + err.Error())
	}
	var setupUrlDefault string
	setupUrl := &setupUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "setup-url", &setupUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "setup-url": ` + err.Error())
	}
	var title string
	err = flags.ParseFlag(cmd.Flags(), "title", &title)
	if err != nil {
		return fmt.Errorf(`error parsing "title": ` + err.Error())
	}
	var userPermissionsFilter []string
	err = flags.ParseFlag(cmd.Flags(), "user-permissions-filter", &userPermissionsFilter)
	if err != nil {
		return fmt.Errorf(`error parsing "user-permissions-filter": ` + err.Error())
	}
	var webhookUrlDefault string
	webhookUrl := &webhookUrlDefault
	err = flags.ParseFlag(cmd.Flags(), "webhook-url", &webhookUrl)
	if err != nil {
		return fmt.Errorf(`error parsing "webhook-url": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.WebAppPut{

		AppPrincipalPermissions: appPrincipalPermissions,
		Description:             description,
		LoginUrl:                loginUrl,
		LogoUrl:                 logoUrl,
		RedirectUrls:            redirectUrls,
		SetupUrl:                setupUrl,
		Title:                   title,
		UserPermissionsFilter:   userPermissionsFilter,
		WebhookUrl:              webhookUrl,
	}

	resp, err := client.AppRegistryService.UpdateApp(appName, model.MakeUpdateAppRequestFromWebAppPut(generated_request_body))
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
