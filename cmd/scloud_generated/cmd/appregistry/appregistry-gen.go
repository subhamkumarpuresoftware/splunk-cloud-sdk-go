// Package appregistry -- generated by scloudgen
// !! DO NOT EDIT !!
//
package appregistry

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud_generated/pkg/appregistry"
)

// createAppNativeApp -- Creates an app.
var createAppNativeAppCmd = &cobra.Command{
	Use:   "create-app-native-app",
	Short: "Creates an app.",
	RunE:  impl.CreateAppNativeApp,
}

// createAppServiceApp -- Creates an app.
var createAppServiceAppCmd = &cobra.Command{
	Use:   "create-app-service-app",
	Short: "Creates an app.",
	RunE:  impl.CreateAppServiceApp,
}

// createAppWebApp -- Creates an app.
var createAppWebAppCmd = &cobra.Command{
	Use:   "create-app-web-app",
	Short: "Creates an app.",
	RunE:  impl.CreateAppWebApp,
}

// createSubscription -- Creates a subscription.
var createSubscriptionCmd = &cobra.Command{
	Use:   "create-subscription",
	Short: "Creates a subscription.",
	RunE:  impl.CreateSubscription,
}

// deleteApp -- Removes an app.
var deleteAppCmd = &cobra.Command{
	Use:   "delete-app",
	Short: "Removes an app.",
	RunE:  impl.DeleteApp,
}

// deleteSubscription -- Removes a subscription.
var deleteSubscriptionCmd = &cobra.Command{
	Use:   "delete-subscription",
	Short: "Removes a subscription.",
	RunE:  impl.DeleteSubscription,
}

// getApp -- Returns the metadata of an app.
var getAppCmd = &cobra.Command{
	Use:   "get-app",
	Short: "Returns the metadata of an app.",
	RunE:  impl.GetApp,
}

// getKeys -- Returns a list of the public keys used for verifying signed webhook requests.
var getKeysCmd = &cobra.Command{
	Use:   "get-keys",
	Short: "Returns a list of the public keys used for verifying signed webhook requests.",
	RunE:  impl.GetKeys,
}

// getSubscription -- Returns or validates a subscription.
var getSubscriptionCmd = &cobra.Command{
	Use:   "get-subscription",
	Short: "Returns or validates a subscription.",
	RunE:  impl.GetSubscription,
}

// listAppSubscriptions -- Returns the collection of subscriptions to an app.
var listAppSubscriptionsCmd = &cobra.Command{
	Use:   "list-app-subscriptions",
	Short: "Returns the collection of subscriptions to an app.",
	RunE:  impl.ListAppSubscriptions,
}

// listApps -- Returns a list of apps.
var listAppsCmd = &cobra.Command{
	Use:   "list-apps",
	Short: "Returns a list of apps.",
	RunE:  impl.ListApps,
}

// listSubscriptions -- Returns the tenant subscriptions.
var listSubscriptionsCmd = &cobra.Command{
	Use:   "list-subscriptions",
	Short: "Returns the tenant subscriptions.",
	RunE:  impl.ListSubscriptions,
}

// rotateSecret -- Rotates the client secret for an app.
var rotateSecretCmd = &cobra.Command{
	Use:   "rotate-secret",
	Short: "Rotates the client secret for an app.",
	RunE:  impl.RotateSecret,
}

// updateApp -- Updates an app.
var updateAppCmd = &cobra.Command{
	Use:   "update-app",
	Short: "Updates an app.",
	RunE:  impl.UpdateApp,
}

func init() {

	appregistryCmd.AddCommand(createAppNativeAppCmd)
	var createAppNativeAppAppPrincipalPermissions []string
	createAppNativeAppCmd.Flags().StringSliceVar(&createAppNativeAppAppPrincipalPermissions, "app-principal-permissions", nil, "Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.")

	var createAppNativeAppDescription string
	createAppNativeAppCmd.Flags().StringVar(&createAppNativeAppDescription, "description", "", "Short paragraph describing the app.")

	var createAppNativeAppKind string
	createAppNativeAppCmd.Flags().StringVar(&createAppNativeAppKind, "kind", "", "This is a required parameter. This is a required parameter.  can accept values web, native, service")
	createAppNativeAppCmd.MarkFlagRequired("kind")

	var createAppNativeAppLoginUrl string
	createAppNativeAppCmd.Flags().StringVar(&createAppNativeAppLoginUrl, "login-url", "", "The URL used to log in to the app.")

	var createAppNativeAppLogoUrl string
	createAppNativeAppCmd.Flags().StringVar(&createAppNativeAppLogoUrl, "logo-url", "", "The URL used to display the app's logo.")

	var createAppNativeAppName string
	createAppNativeAppCmd.Flags().StringVar(&createAppNativeAppName, "name", "", "This is a required parameter. This is a required parameter. App name that is unique within Splunk Cloud Platform.")
	createAppNativeAppCmd.MarkFlagRequired("name")

	var createAppNativeAppRedirectUrls []string
	createAppNativeAppCmd.Flags().StringSliceVar(&createAppNativeAppRedirectUrls, "redirect-urls", nil, "Array of URLs that can be used for redirect after logging into the app.")

	var createAppNativeAppSetupUrl string
	createAppNativeAppCmd.Flags().StringVar(&createAppNativeAppSetupUrl, "setup-url", "", "URL to redirect to after a subscription is created.")

	var createAppNativeAppTitle string
	createAppNativeAppCmd.Flags().StringVar(&createAppNativeAppTitle, "title", "", "This is a required parameter. Human-readable title for the app.")
	createAppNativeAppCmd.MarkFlagRequired("title")

	var createAppNativeAppUserPermissionsFilter []string
	createAppNativeAppCmd.Flags().StringSliceVar(&createAppNativeAppUserPermissionsFilter, "user-permissions-filter", nil, "Array of permission filter templates that are used to intersect with a user's permissions when using the app.")

	var createAppNativeAppWebhookUrl string
	createAppNativeAppCmd.Flags().StringVar(&createAppNativeAppWebhookUrl, "webhook-url", "", "URL that webhook events are sent to.")

	appregistryCmd.AddCommand(createAppServiceAppCmd)
	var createAppServiceAppAppPrincipalPermissions []string
	createAppServiceAppCmd.Flags().StringSliceVar(&createAppServiceAppAppPrincipalPermissions, "app-principal-permissions", nil, "Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.")

	var createAppServiceAppDescription string
	createAppServiceAppCmd.Flags().StringVar(&createAppServiceAppDescription, "description", "", "Short paragraph describing the app.")

	var createAppServiceAppKind string
	createAppServiceAppCmd.Flags().StringVar(&createAppServiceAppKind, "kind", "", "This is a required parameter. This is a required parameter.  can accept values web, native, service")
	createAppServiceAppCmd.MarkFlagRequired("kind")

	var createAppServiceAppLoginUrl string
	createAppServiceAppCmd.Flags().StringVar(&createAppServiceAppLoginUrl, "login-url", "", "The URL used to log in to the app.")

	var createAppServiceAppLogoUrl string
	createAppServiceAppCmd.Flags().StringVar(&createAppServiceAppLogoUrl, "logo-url", "", "The URL used to display the app's logo.")

	var createAppServiceAppName string
	createAppServiceAppCmd.Flags().StringVar(&createAppServiceAppName, "name", "", "This is a required parameter. This is a required parameter. App name that is unique within Splunk Cloud Platform.")
	createAppServiceAppCmd.MarkFlagRequired("name")

	var createAppServiceAppRedirectUrls []string
	createAppServiceAppCmd.Flags().StringSliceVar(&createAppServiceAppRedirectUrls, "redirect-urls", nil, "Array of URLs that can be used for redirect after logging into the app.")

	var createAppServiceAppSetupUrl string
	createAppServiceAppCmd.Flags().StringVar(&createAppServiceAppSetupUrl, "setup-url", "", "URL to redirect to after a subscription is created.")

	var createAppServiceAppTitle string
	createAppServiceAppCmd.Flags().StringVar(&createAppServiceAppTitle, "title", "", "This is a required parameter. Human-readable title for the app.")
	createAppServiceAppCmd.MarkFlagRequired("title")

	var createAppServiceAppUserPermissionsFilter []string
	createAppServiceAppCmd.Flags().StringSliceVar(&createAppServiceAppUserPermissionsFilter, "user-permissions-filter", nil, "Array of permission filter templates that are used to intersect with a user's permissions when using the app.")

	var createAppServiceAppWebhookUrl string
	createAppServiceAppCmd.Flags().StringVar(&createAppServiceAppWebhookUrl, "webhook-url", "", "URL that webhook events are sent to.")

	appregistryCmd.AddCommand(createAppWebAppCmd)
	var createAppWebAppAppPrincipalPermissions []string
	createAppWebAppCmd.Flags().StringSliceVar(&createAppWebAppAppPrincipalPermissions, "app-principal-permissions", nil, "Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.")

	var createAppWebAppDescription string
	createAppWebAppCmd.Flags().StringVar(&createAppWebAppDescription, "description", "", "Short paragraph describing the app.")

	var createAppWebAppKind string
	createAppWebAppCmd.Flags().StringVar(&createAppWebAppKind, "kind", "", "This is a required parameter. This is a required parameter.  can accept values web, native, service")
	createAppWebAppCmd.MarkFlagRequired("kind")

	var createAppWebAppLoginUrl string
	createAppWebAppCmd.Flags().StringVar(&createAppWebAppLoginUrl, "login-url", "", "The URL used to log in to the app.")

	var createAppWebAppLogoUrl string
	createAppWebAppCmd.Flags().StringVar(&createAppWebAppLogoUrl, "logo-url", "", "The URL used to display the app's logo.")

	var createAppWebAppName string
	createAppWebAppCmd.Flags().StringVar(&createAppWebAppName, "name", "", "This is a required parameter. This is a required parameter. App name that is unique within Splunk Cloud Platform.")
	createAppWebAppCmd.MarkFlagRequired("name")

	var createAppWebAppRedirectUrls []string
	createAppWebAppCmd.Flags().StringSliceVar(&createAppWebAppRedirectUrls, "redirect-urls", nil, "This is a required parameter. Array of URLs that can be used for redirect after logging into the app.")
	createAppWebAppCmd.MarkFlagRequired("redirectUrls")

	var createAppWebAppSetupUrl string
	createAppWebAppCmd.Flags().StringVar(&createAppWebAppSetupUrl, "setup-url", "", "URL to redirect to after a subscription is created.")

	var createAppWebAppTitle string
	createAppWebAppCmd.Flags().StringVar(&createAppWebAppTitle, "title", "", "This is a required parameter. Human-readable title for the app.")
	createAppWebAppCmd.MarkFlagRequired("title")

	var createAppWebAppUserPermissionsFilter []string
	createAppWebAppCmd.Flags().StringSliceVar(&createAppWebAppUserPermissionsFilter, "user-permissions-filter", nil, "Array of permission filter templates that are used to intersect with a user's permissions when using the app.")

	var createAppWebAppWebhookUrl string
	createAppWebAppCmd.Flags().StringVar(&createAppWebAppWebhookUrl, "webhook-url", "", "URL that webhook events are sent to.")

	appregistryCmd.AddCommand(createSubscriptionCmd)

	var createSubscriptionAppName string
	createSubscriptionCmd.Flags().StringVar(&createSubscriptionAppName, "app-name", "", "This is a required parameter. ")
	createSubscriptionCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(deleteAppCmd)

	var deleteAppAppName string
	deleteAppCmd.Flags().StringVar(&deleteAppAppName, "app-name", "", "This is a required parameter. App name.")
	deleteAppCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(deleteSubscriptionCmd)

	var deleteSubscriptionAppName string
	deleteSubscriptionCmd.Flags().StringVar(&deleteSubscriptionAppName, "app-name", "", "This is a required parameter. App name.")
	deleteSubscriptionCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(getAppCmd)

	var getAppAppName string
	getAppCmd.Flags().StringVar(&getAppAppName, "app-name", "", "This is a required parameter. App name.")
	getAppCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(getKeysCmd)

	appregistryCmd.AddCommand(getSubscriptionCmd)

	var getSubscriptionAppName string
	getSubscriptionCmd.Flags().StringVar(&getSubscriptionAppName, "app-name", "", "This is a required parameter. App name.")
	getSubscriptionCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(listAppSubscriptionsCmd)

	var listAppSubscriptionsAppName string
	listAppSubscriptionsCmd.Flags().StringVar(&listAppSubscriptionsAppName, "app-name", "", "This is a required parameter. App name.")
	listAppSubscriptionsCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(listAppsCmd)

	appregistryCmd.AddCommand(listSubscriptionsCmd)

	var listSubscriptionsKind string
	listSubscriptionsCmd.Flags().StringVar(&listSubscriptionsKind, "kind", "", "The kind of application.")

	appregistryCmd.AddCommand(rotateSecretCmd)

	var rotateSecretAppName string
	rotateSecretCmd.Flags().StringVar(&rotateSecretAppName, "app-name", "", "This is a required parameter. App name.")
	rotateSecretCmd.MarkFlagRequired("app-name")

	appregistryCmd.AddCommand(updateAppCmd)

	var updateAppAppName string
	updateAppCmd.Flags().StringVar(&updateAppAppName, "app-name", "", "This is a required parameter. App name.")
	updateAppCmd.MarkFlagRequired("app-name")

	var updateAppAppPrincipalPermissions []string
	updateAppCmd.Flags().StringSliceVar(&updateAppAppPrincipalPermissions, "app-principal-permissions", nil, "Array of permission templates that are used to grant permission to the app principal when a tenant subscribes.")

	var updateAppDescription string
	updateAppCmd.Flags().StringVar(&updateAppDescription, "description", "", "Short paragraph describing the app.")

	var updateAppLoginUrl string
	updateAppCmd.Flags().StringVar(&updateAppLoginUrl, "login-url", "", "The URL used to log in to the app.")

	var updateAppLogoUrl string
	updateAppCmd.Flags().StringVar(&updateAppLogoUrl, "logo-url", "", "The URL used to display the app's logo.")

	var updateAppRedirectUrls []string
	updateAppCmd.Flags().StringSliceVar(&updateAppRedirectUrls, "redirect-urls", nil, "Array of URLs that can be used for redirect after logging into the app.")

	var updateAppSetupUrl string
	updateAppCmd.Flags().StringVar(&updateAppSetupUrl, "setup-url", "", "URL to redirect to after a subscription is created.")

	var updateAppTitle string
	updateAppCmd.Flags().StringVar(&updateAppTitle, "title", "", "Human-readable title for the app.")

	var updateAppUserPermissionsFilter []string
	updateAppCmd.Flags().StringSliceVar(&updateAppUserPermissionsFilter, "user-permissions-filter", nil, "Array of permission filter templates that are used to intersect with a user's permissions when using the app.")

	var updateAppWebhookUrl string
	updateAppCmd.Flags().StringVar(&updateAppWebhookUrl, "webhook-url", "", "URL that webhook events are sent to.")

}
