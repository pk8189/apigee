
## SDK Usage Examples


### 
Deletes a category from the portal.

**API Endpoint**: `DELETE /v1/{name}`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	v1 "apigee_api/resources/v1"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Delete(v1.DeleteRequest { Name: "string" })
```

    
### 
Deletes the data from a debug session. This does not cancel the debug session or prevent further data from being collected if the session is still active in runtime pods.

**API Endpoint**: `DELETE /v1/{name}/data`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	data "apigee_api/resources/v1/data"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Data.Delete(data.DeleteRequest { Name: "string" })
```

    
### 
Undeploys a shared flow revision from an environment. For a request path `organizations/{org}/environments/{env}/sharedflows/{sf}/revisions/{rev}/deployments`, two permissions are required: * `apigee.deployments.delete` on the resource `organizations/{org}/environments/{env}` * `apigee.sharedflowrevisions.undeploy` on the resource `organizations/{org}/sharedflows/{sf}/revisions/{rev}`

**API Endpoint**: `DELETE /v1/{name}/deployments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	deployments "apigee_api/resources/v1/deployments"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Deployments.Delete(deployments.DeleteRequest { Name: "string" })
```

    
### 
Deletes a resource file. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).

**API Endpoint**: `DELETE /v1/{parent}/resourcefiles/{type}/{name}`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	x "apigee_api/resources/v1/resourcefiles/x"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Resourcefiles.X.Delete(x.DeleteRequest { Parent: "string", Type: "string", Name: "string" })
```

    
### 
Gets a category on the portal.

**API Endpoint**: `GET /v1/{name}`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	v1 "apigee_api/resources/v1"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.GetByName(v1.GetByNameRequest { Name: "string" })
```

    
### 
Gets the certificate from an alias in PEM-encoded form.

**API Endpoint**: `GET /v1/{name}/certificate`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	certificate "apigee_api/resources/v1/certificate"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Certificate.List(certificate.ListRequest { Name: "string" })
```

    
### 
Generates a PKCS #10 Certificate Signing Request for the private key in an alias.

**API Endpoint**: `GET /v1/{name}/csr`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	csr "apigee_api/resources/v1/csr"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Csr.List(csr.ListRequest { Name: "string" })
```

    
### 
Gets the deployment of a shared flow revision and actual state reported by runtime pods.

**API Endpoint**: `GET /v1/{name}/deployments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	deployments "apigee_api/resources/v1/deployments"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Deployments.List1(deployments.List1Request { Name: "string" })
```

    
### 
Lists operations that match the specified filter in the request. If the server doesn't support this method, it returns `UNIMPLEMENTED`.

**API Endpoint**: `GET /v1/{name}/operations`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	operations "apigee_api/resources/v1/operations"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Operations.List(operations.ListRequest { Name: "string" })
```

    
### 
Gets the project ID and region for an Apigee organization.

**API Endpoint**: `GET /v1/{name}:getProjectMapping`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	get_project_mapping "apigee_api/resources/v1/get_project_mapping"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.GetProjectMapping.List(get_project_mapping.ListRequest { Name: "string" })
```

    
### 
ListSecurityProfileRevisions lists all the revisions of the security profile.

**API Endpoint**: `GET /v1/{name}:listRevisions`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	list_revisions "apigee_api/resources/v1/list_revisions"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.ListRevisions.List(list_revisions.ListRequest { Name: "string" })
```

    
### 
Lists the Apigee organizations and associated Google Cloud projects that you have permission to access. See [Understanding organizations](https://cloud.google.com/apigee/docs/api-platform/fundamentals/organization-structure).

**API Endpoint**: `GET /v1/{parent}`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	v1 "apigee_api/resources/v1"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.GetByParent(v1.GetByParentRequest { Parent: "string" })
```

    
### 
List Datastores

**API Endpoint**: `GET /v1/{parent}/analytics/datastores`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	datastores "apigee_api/resources/v1/analytics/datastores"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Analytics.Datastores.List(datastores.ListRequest { Parent: "string" })
```

    
### 
Lists the details and status of all analytics export jobs belonging to the parent organization and environment.

**API Endpoint**: `GET /v1/{parent}/analytics/exports`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	exports "apigee_api/resources/v1/analytics/exports"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Analytics.Exports.List(exports.ListRequest { Parent: "string" })
```

    
### 
Lists the categories on the portal.

**API Endpoint**: `GET /v1/{parent}/apicategories`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	apicategories "apigee_api/resources/v1/apicategories"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Apicategories.List(apicategories.ListRequest { Parent: "string" })
```

    
### 
Lists all API product names for an organization. Filter the list by passing an `attributename` and `attibutevalue`. The maximum number of API products returned is 1000. You can paginate the list of API products returned using the `startKey` and `count` query parameters.

**API Endpoint**: `GET /v1/{parent}/apiproducts`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	apiproducts "apigee_api/resources/v1/apiproducts"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Apiproducts.List(apiproducts.ListRequest { Parent: "string" })
```

    
### 
Lists the names of all API proxies in an organization. The names returned correspond to the names defined in the configuration files for each API proxy.

**API Endpoint**: `GET /v1/{parent}/apis`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	apis "apigee_api/resources/v1/apis"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Apis.List(apis.ListRequest { Parent: "string" })
```

    
### 
Lists all apps created by a developer in an Apigee organization. Optionally, you can request an expanded view of the developer apps. A maximum of 100 developer apps are returned per API call. You can paginate the list of deveoper apps returned using the `startKey` and `count` query parameters.

**API Endpoint**: `GET /v1/{parent}/apps`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	apps "apigee_api/resources/v1/apps"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Apps.List(apps.ListRequest { Parent: "string" })
```

    
### 
Lists the ArchiveDeployments in the specified Environment.

**API Endpoint**: `GET /v1/{parent}/archiveDeployments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	archive_deployments "apigee_api/resources/v1/archive_deployments"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.ArchiveDeployments.List(archive_deployments.ListRequest { Parent: "string" })
```

    
### 
Lists all attachments to an instance. **Note:** Not supported for Apigee hybrid.

**API Endpoint**: `GET /v1/{parent}/attachments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	attachments "apigee_api/resources/v1/attachments"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Attachments.List(attachments.ListRequest { Parent: "string" })
```

    
### 
Returns a list of all developer attributes.

**API Endpoint**: `GET /v1/{parent}/attributes`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	attributes "apigee_api/resources/v1/attributes"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Attributes.List(attributes.ListRequest { Parent: "string" })
```

    
### 
Lists all data collectors.

**API Endpoint**: `GET /v1/{parent}/datacollectors`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	datacollectors "apigee_api/resources/v1/datacollectors"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Datacollectors.List(datacollectors.ListRequest { Parent: "string" })
```

    
### 
Lists debug sessions that are currently active in the given API Proxy revision.

**API Endpoint**: `GET /v1/{parent}/debugsessions`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	debugsessions "apigee_api/resources/v1/debugsessions"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Debugsessions.List(debugsessions.ListRequest { Parent: "string" })
```

    
### 
Lists all deployments of a shared flow revision.

**API Endpoint**: `GET /v1/{parent}/deployments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	deployments "apigee_api/resources/v1/deployments"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Deployments.List1(deployments.List1Request { Parent: "string" })
```

    
### 
Lists all developers in an organization by email address. By default, the response does not include company developers. Set the `includeCompany` query parameter to `true` to include company developers. **Note**: A maximum of 1000 developers are returned in the response. You paginate the list of developers returned using the `startKey` and `count` query parameters.

**API Endpoint**: `GET /v1/{parent}/developers`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	developers "apigee_api/resources/v1/developers"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Developers.List(developers.ListRequest { Parent: "string" })
```

    
### 
Lists the endpoint attachments in an organization.

**API Endpoint**: `GET /v1/{parent}/endpointAttachments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	endpoint_attachments "apigee_api/resources/v1/endpoint_attachments"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.EndpointAttachments.List(endpoint_attachments.ListRequest { Parent: "string" })
```

    
### 
Lists key value entries for key values maps scoped to an organization, environment, or API proxy. **Note**: Supported for Apigee hybrid 1.8.x and higher.

**API Endpoint**: `GET /v1/{parent}/entries`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	entries "apigee_api/resources/v1/entries"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Entries.List(entries.ListRequest { Parent: "string" })
```

    
### 
Lists all environment groups.

**API Endpoint**: `GET /v1/{parent}/envgroups`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	envgroups "apigee_api/resources/v1/envgroups"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Envgroups.List(envgroups.ListRequest { Parent: "string" })
```

    
### 
Return a list of Asynchronous Queries at host level.

**API Endpoint**: `GET /v1/{parent}/hostQueries`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	host_queries "apigee_api/resources/v1/host_queries"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.HostQueries.List(host_queries.ListRequest { Parent: "string" })
```

    
### 
Return a list of Security Reports at host level.

**API Endpoint**: `GET /v1/{parent}/hostSecurityReports`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	host_security_reports "apigee_api/resources/v1/host_security_reports"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.HostSecurityReports.List(host_security_reports.ListRequest { Parent: "string" })
```

    
### 
Lists all Apigee runtime instances for the organization. **Note:** Not supported for Apigee hybrid.

**API Endpoint**: `GET /v1/{parent}/instances`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	instances "apigee_api/resources/v1/instances"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Instances.List(instances.ListRequest { Parent: "string" })
```

    
### 
Lists the NAT addresses for an Apigee instance. **Note:** Not supported for Apigee hybrid.

**API Endpoint**: `GET /v1/{parent}/natAddresses`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	nat_addresses "apigee_api/resources/v1/nat_addresses"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.NatAddresses.List(nat_addresses.ListRequest { Parent: "string" })
```

    
### 
Lists all of the distributed trace configuration overrides in an environment.

**API Endpoint**: `GET /v1/{parent}/overrides`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	overrides "apigee_api/resources/v1/overrides"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Overrides.List(overrides.ListRequest { Parent: "string" })
```

    
### 
Return a list of Asynchronous Queries

**API Endpoint**: `GET /v1/{parent}/queries`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	queries "apigee_api/resources/v1/queries"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Queries.List(queries.ListRequest { Parent: "string" })
```

    
### 
Lists all the rate plans for an API product.

**API Endpoint**: `GET /v1/{parent}/rateplans`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	rateplans "apigee_api/resources/v1/rateplans"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Rateplans.List(rateplans.ListRequest { Parent: "string" })
```

    
### 
Return a list of Custom Reports

**API Endpoint**: `GET /v1/{parent}/reports`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	reports "apigee_api/resources/v1/reports"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Reports.List(reports.ListRequest { Parent: "string" })
```

    
### 
Lists all resource files, optionally filtering by type. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).

**API Endpoint**: `GET /v1/{parent}/resourcefiles`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	resourcefiles "apigee_api/resources/v1/resourcefiles"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Resourcefiles.List(resourcefiles.ListRequest { Parent: "string" })
```

    
### 
Lists all resource files, optionally filtering by type. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).

**API Endpoint**: `GET /v1/{parent}/resourcefiles/{type}`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	resourcefiles "apigee_api/resources/v1/resourcefiles"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Resourcefiles.Get(resourcefiles.GetRequest { Parent: "string", Type: "string" })
```

    
### 
Gets the contents of a resource file. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).

**API Endpoint**: `GET /v1/{parent}/resourcefiles/{type}/{name}`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	x "apigee_api/resources/v1/resourcefiles/x"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Resourcefiles.X.Get(x.GetRequest { Parent: "string", Type: "string", Name: "string" })
```

    
### 
ListSecurityIncidents lists all the security incident associated with the environment.

**API Endpoint**: `GET /v1/{parent}/securityIncidents`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	security_incidents "apigee_api/resources/v1/security_incidents"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.SecurityIncidents.List(security_incidents.ListRequest { Parent: "string" })
```

    
### 
ListSecurityProfiles lists all the security profiles associated with the org including attached and unattached profiles.

**API Endpoint**: `GET /v1/{parent}/securityProfiles`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	security_profiles "apigee_api/resources/v1/security_profiles"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.SecurityProfiles.List(security_profiles.ListRequest { Parent: "string" })
```

    
### 
Return a list of Security Reports

**API Endpoint**: `GET /v1/{parent}/securityReports`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	security_reports "apigee_api/resources/v1/security_reports"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.SecurityReports.List(security_reports.ListRequest { Parent: "string" })
```

    
### 
Lists all shared flows in the organization.

**API Endpoint**: `GET /v1/{parent}/sharedflows`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	sharedflows "apigee_api/resources/v1/sharedflows"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Sharedflows.List(sharedflows.ListRequest { Parent: "string" })
```

    
### 
Lists all API product subscriptions for a developer.

**API Endpoint**: `GET /v1/{parent}/subscriptions`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	subscriptions "apigee_api/resources/v1/subscriptions"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Subscriptions.List(subscriptions.ListRequest { Parent: "string" })
```

    
### 
Gets the IAM policy on an environment. For more information, see [Manage users, roles, and permissions using the API](https://cloud.google.com/apigee/docs/api-platform/system-administration/manage-users-roles). You must have the `apigee.environments.getIamPolicy` permission to call this API.

**API Endpoint**: `GET /v1/{resource}:getIamPolicy`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	get_iam_policy "apigee_api/resources/v1/get_iam_policy"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.GetIamPolicy.List(get_iam_policy.ListRequest { Resource: "string" })
```

    
### 
Updates a category on the portal.

**API Endpoint**: `PATCH /v1/{name}`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	v1 "apigee_api/resources/v1"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Patch(v1.PatchRequest { Name: "string", Data: types.GoogleCloudApigeeV1ApiCategoryData { Id: nullable.NewValue("string"), Name: nullable.NewValue("string"), SiteId: nullable.NewValue("string"), UpdateTime: nullable.NewValue("string") } })
```

    
### 
Creates an Apigee organization. See [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).

**API Endpoint**: `POST /v1/organizations`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	organizations "apigee_api/resources/v1/organizations"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Organizations.Create(organizations.CreateRequest { Data: types.GoogleCloudApigeeV1Organization { AddonsConfig: nullable.NewValue(types.GoogleCloudApigeeV1AddonsConfig { AdvancedApiOpsConfig: nullable.NewValue(types.GoogleCloudApigeeV1AdvancedApiOpsConfig { Enabled: nullable.NewValue(true) }), ApiSecurityConfig: nullable.NewValue(types.GoogleCloudApigeeV1ApiSecurityConfig { Enabled: nullable.NewValue(true), ExpiresAt: nullable.NewValue("string") }), ConnectorsPlatformConfig: nullable.NewValue(types.GoogleCloudApigeeV1ConnectorsPlatformConfig { Enabled: nullable.NewValue(true), ExpiresAt: nullable.NewValue("string") }), IntegrationConfig: nullable.NewValue(types.GoogleCloudApigeeV1IntegrationConfig { Enabled: nullable.NewValue(true) }), MonetizationConfig: nullable.NewValue(types.GoogleCloudApigeeV1MonetizationConfig { Enabled: nullable.NewValue(true) }) }), AnalyticsRegion: nullable.NewValue("string"), ApiConsumerDataEncryptionKeyName: nullable.NewValue("string"), ApiConsumerDataLocation: nullable.NewValue("string"), ApigeeProjectId: nullable.NewValue("string"), Attributes: nullable.NewValue([]string{"string"}), AuthorizedNetwork: nullable.NewValue("string"), BillingType: nullable.NewValue(types.GoogleCloudApigeeV1OrganizationBillingTypeEnumBillingTypeUnspecified), CaCertificate: nullable.NewValue("string"), ControlPlaneEncryptionKeyName: nullable.NewValue("string"), CreatedAt: nullable.NewValue("string"), CustomerName: nullable.NewValue("string"), Description: nullable.NewValue("string"), DisplayName: nullable.NewValue("string"), Environments: nullable.NewValue([]string{"string"}), ExpiresAt: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), PortalDisabled: nullable.NewValue(true), ProjectId: nullable.NewValue("string"), Properties: nullable.NewValue(types.GoogleCloudApigeeV1Properties { Property: nullable.NewValue([]types.GoogleCloudApigeeV1Property{types.GoogleCloudApigeeV1Property { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}) }), RuntimeDatabaseEncryptionKeyName: nullable.NewValue("string"), RuntimeType: nullable.NewValue(types.GoogleCloudApigeeV1OrganizationRuntimeTypeEnumCloud), State: nullable.NewValue(types.GoogleCloudApigeeV1OrganizationStateEnumActive), SubscriptionType: nullable.NewValue(types.GoogleCloudApigeeV1OrganizationSubscriptionTypeEnumPaid), Type: nullable.NewValue(types.GoogleCloudApigeeV1OrganizationTypeEnumTypeInternal) } })
```

    
### 
Reports the latest status for a runtime instance.

**API Endpoint**: `POST /v1/{instance}:reportStatus`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	report_status "apigee_api/resources/v1/report_status"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.ReportStatus.Create(report_status.CreateRequest { Instance: "string", Data: types.GoogleCloudApigeeV1ReportInstanceStatusRequest { InstanceUid: nullable.NewValue("string"), ReportTime: nullable.NewValue("string"), Resources: nullable.NewValue([]types.GoogleCloudApigeeV1ResourceStatus{types.GoogleCloudApigeeV1ResourceStatus { Resource: nullable.NewValue("string"), Revisions: nullable.NewValue([]types.GoogleCloudApigeeV1RevisionStatus{types.GoogleCloudApigeeV1RevisionStatus { Errors: nullable.NewValue([]types.GoogleCloudApigeeV1UpdateError{types.GoogleCloudApigeeV1UpdateError { Code: nullable.NewValue(types.GoogleCloudApigeeV1UpdateErrorCodeEnumAborted), Message: nullable.NewValue("string"), Resource: nullable.NewValue("string"), Type: nullable.NewValue("string") }}), JsonSpec: nullable.NewValue("string"), Replicas: nullable.NewValue(123), RevisionId: nullable.NewValue("string") }}), TotalReplicas: nullable.NewValue(123), Uid: nullable.NewValue("string") }}) } })
```

    
### 
Updates a shared flow revision. This operation is only allowed on revisions which have never been deployed. After deployment a revision becomes immutable, even if it becomes undeployed. The payload is a ZIP-formatted shared flow. Content type must be either multipart/form-data or application/octet-stream.

**API Endpoint**: `POST /v1/{name}`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	v1 "apigee_api/resources/v1"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Create(v1.CreateRequest { Name: "string", Data: types.GoogleApiHttpBody { ContentType: nullable.NewValue("string"), Data: nullable.NewValue("string"), Extensions: nullable.NewValue([]map[string]interface{}{map[string]interface{}{}}) } })
```

    
### 
Updates attributes for a developer app. This API replaces the current attributes with those specified in the request.

**API Endpoint**: `POST /v1/{name}/attributes`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	attributes "apigee_api/resources/v1/attributes"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Attributes.Create(attributes.CreateRequest { Name: "string", Data: types.GoogleCloudApigeeV1Attributes { Attribute: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}) } })
```

    
### 
Deploys a revision of a shared flow. If another revision of the same shared flow is currently deployed, set the `override` parameter to `true` to have this revision replace the currently deployed revision. You cannot use a shared flow until it has been deployed to an environment. For a request path `organizations/{org}/environments/{env}/sharedflows/{sf}/revisions/{rev}/deployments`, two permissions are required: * `apigee.deployments.create` on the resource `organizations/{org}/environments/{env}` * `apigee.sharedflowrevisions.deploy` on the resource `organizations/{org}/sharedflows/{sf}/revisions/{rev}`

**API Endpoint**: `POST /v1/{name}/deployments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	deployments "apigee_api/resources/v1/deployments"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Deployments.Create(deployments.CreateRequest { Name: "string" })
```

    
### 
Generates a report for a dry run analysis of a DeployApiProxy request without committing the deployment. In addition to the standard validations performed when adding deployments, additional analysis will be done to detect possible traffic routing changes that would result from this deployment being created. Any potential routing conflicts or unsafe changes will be reported in the response. This routing analysis is not performed for a non-dry-run DeployApiProxy request. For a request path `organizations/{org}/environments/{env}/apis/{api}/revisions/{rev}/deployments:generateDeployChangeReport`, two permissions are required: * `apigee.deployments.create` on the resource `organizations/{org}/environments/{env}` * `apigee.proxyrevisions.deploy` on the resource `organizations/{org}/apis/{api}/revisions/{rev}`

**API Endpoint**: `POST /v1/{name}/deployments:generateDeployChangeReport`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	deployments_generate_deploy_change_report "apigee_api/resources/v1/deployments_generate_deploy_change_report"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.DeploymentsGenerateDeployChangeReport.Create(deployments_generate_deploy_change_report.CreateRequest { Name: "string" })
```

    
### 
Generates a report for a dry run analysis of an UndeployApiProxy request without committing the undeploy. In addition to the standard validations performed when removing deployments, additional analysis will be done to detect possible traffic routing changes that would result from this deployment being removed. Any potential routing conflicts or unsafe changes will be reported in the response. This routing analysis is not performed for a non-dry-run UndeployApiProxy request. For a request path `organizations/{org}/environments/{env}/apis/{api}/revisions/{rev}/deployments:generateUndeployChangeReport`, two permissions are required: * `apigee.deployments.delete` on the resource `organizations/{org}/environments/{env}` * `apigee.proxyrevisions.undeploy` on the resource `organizations/{org}/apis/{api}/revisions/{rev}`

**API Endpoint**: `POST /v1/{name}/deployments:generateUndeployChangeReport`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	deployments_generate_undeploy_change_report "apigee_api/resources/v1/deployments_generate_undeploy_change_report"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.DeploymentsGenerateUndeployChangeReport.Create(deployments_generate_undeploy_change_report.CreateRequest { Name: "string" })
```

    
### 
Activates the NAT address. The Apigee instance can now use this for Internet egress traffic. **Note:** Not supported for Apigee hybrid.

**API Endpoint**: `POST /v1/{name}:activate`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	activate "apigee_api/resources/v1/activate"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Activate.Create(activate.CreateRequest { Name: "string", Data: map[string]interface{}{} })
```

    
### 
Adjust the prepaid balance for the developer. This API will be used in scenarios where the developer has been under-charged or over-charged.

**API Endpoint**: `POST /v1/{name}:adjust`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	adjust "apigee_api/resources/v1/adjust"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Adjust.Create(adjust.CreateRequest { Name: "string", Data: types.GoogleCloudApigeeV1AdjustDeveloperBalanceRequest { Adjustment: nullable.NewValue(types.GoogleTypeMoney { CurrencyCode: nullable.NewValue("string"), Nanos: nullable.NewValue(123), Units: nullable.NewValue("string") }) } })
```

    
### 
Credits the account balance for the developer.

**API Endpoint**: `POST /v1/{name}:credit`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	credit "apigee_api/resources/v1/credit"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Credit.Create(credit.CreateRequest { Name: "string", Data: types.GoogleCloudApigeeV1CreditDeveloperBalanceRequest { TransactionAmount: nullable.NewValue(types.GoogleTypeMoney { CurrencyCode: nullable.NewValue("string"), Nanos: nullable.NewValue(123), Units: nullable.NewValue("string") }), TransactionId: nullable.NewValue("string") } })
```

    
### 
Expires an API product subscription immediately.

**API Endpoint**: `POST /v1/{name}:expire`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	expire "apigee_api/resources/v1/expire"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Expire.Create(expire.CreateRequest { Name: "string", Data: map[string]interface{}{} })
```

    
### 
Generates a signed URL for downloading the original zip file used to create an Archive Deployment. The URL is only valid for a limited period and should be used within minutes after generation. Each call returns a new upload URL.

**API Endpoint**: `POST /v1/{name}:generateDownloadUrl`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	generate_download_url "apigee_api/resources/v1/generate_download_url"
	url "net/url"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.GenerateDownloadUrl.Create(generate_download_url.CreateRequest { Name: "string", Data: map[string]interface{}{} })
```

    
### 
Lists the service accounts with the permissions required to allow the Synchronizer to download environment data from the control plane. An ETag is returned in the response to `getSyncAuthorization`. Pass that ETag when calling [setSyncAuthorization](setSyncAuthorization) to ensure that you are updating the correct version. If you don't pass the ETag in the call to `setSyncAuthorization`, then the existing authorization is overwritten indiscriminately. For more information, see [Configure the Synchronizer](https://cloud.google.com/apigee/docs/hybrid/latest/synchronizer-access). **Note**: Available to Apigee hybrid only.

**API Endpoint**: `POST /v1/{name}:getSyncAuthorization`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	get_sync_authorization "apigee_api/resources/v1/get_sync_authorization"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.GetSyncAuthorization.Create(get_sync_authorization.CreateRequest { Name: "string", Data: map[string]interface{}{} })
```

    
### 
Sets the permissions required to allow the Synchronizer to download environment data from the control plane. You must call this API to enable proper functioning of hybrid. Pass the ETag when calling `setSyncAuthorization` to ensure that you are updating the correct version. To get an ETag, call [getSyncAuthorization](getSyncAuthorization). If you don't pass the ETag in the call to `setSyncAuthorization`, then the existing authorization is overwritten indiscriminately. For more information, see [Configure the Synchronizer](https://cloud.google.com/apigee/docs/hybrid/latest/synchronizer-access). **Note**: Available to Apigee hybrid only.

**API Endpoint**: `POST /v1/{name}:setSyncAuthorization`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	set_sync_authorization "apigee_api/resources/v1/set_sync_authorization"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.SetSyncAuthorization.Create(set_sync_authorization.CreateRequest { Name: "string", Data: types.GoogleCloudApigeeV1SyncAuthorization { Etag: nullable.NewValue("string"), Identities: nullable.NewValue([]string{"string"}) } })
```

    
### 
Retrieve security statistics as tabular rows.

**API Endpoint**: `POST /v1/{orgenv}/securityStats:queryTabularStats`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	security_stats_query_tabular_stats "apigee_api/resources/v1/security_stats_query_tabular_stats"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.SecurityStatsQueryTabularStats.Create(security_stats_query_tabular_stats.CreateRequest { Orgenv: "string", Data: types.GoogleCloudApigeeV1QueryTabularStatsRequest { Dimensions: nullable.NewValue([]string{"string"}), Filter: nullable.NewValue("string"), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1MetricAggregation{types.GoogleCloudApigeeV1MetricAggregation { Aggregation: nullable.NewValue(types.GoogleCloudApigeeV1MetricAggregationAggregationEnumAggregationFunctionUnspecified), Name: nullable.NewValue("string"), Order: nullable.NewValue(types.GoogleCloudApigeeV1MetricAggregationOrderEnumAscending) }}), PageSize: nullable.NewValue(123), PageToken: nullable.NewValue("string"), TimeRange: nullable.NewValue(types.GoogleTypeInterval { EndTime: nullable.NewValue("string"), StartTime: nullable.NewValue("string") }) } })
```

    
### 
Retrieve security statistics as a collection of time series.

**API Endpoint**: `POST /v1/{orgenv}/securityStats:queryTimeSeriesStats`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	security_stats_query_time_series_stats "apigee_api/resources/v1/security_stats_query_time_series_stats"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.SecurityStatsQueryTimeSeriesStats.Create(security_stats_query_time_series_stats.CreateRequest { Orgenv: "string", Data: types.GoogleCloudApigeeV1QueryTimeSeriesStatsRequest { Dimensions: nullable.NewValue([]string{"string"}), Filter: nullable.NewValue("string"), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1MetricAggregation{types.GoogleCloudApigeeV1MetricAggregation { Aggregation: nullable.NewValue(types.GoogleCloudApigeeV1MetricAggregationAggregationEnumAggregationFunctionUnspecified), Name: nullable.NewValue("string"), Order: nullable.NewValue(types.GoogleCloudApigeeV1MetricAggregationOrderEnumAscending) }}), PageSize: nullable.NewValue(123), PageToken: nullable.NewValue("string"), TimeRange: nullable.NewValue(types.GoogleTypeInterval { EndTime: nullable.NewValue("string"), StartTime: nullable.NewValue("string") }), TimestampOrder: nullable.NewValue(types.GoogleCloudApigeeV1QueryTimeSeriesStatsRequestTimestampOrderEnumAscending), WindowSize: nullable.NewValue(types.GoogleCloudApigeeV1QueryTimeSeriesStatsRequestWindowSizeEnumDay) } })
```

    
### 
Configures the add-ons for the Apigee organization. The existing add-on configuration will be fully replaced.

**API Endpoint**: `POST /v1/{org}:setAddons`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	set_addons "apigee_api/resources/v1/set_addons"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.SetAddons.Create(set_addons.CreateRequest { Org: "string", Data: types.GoogleCloudApigeeV1SetAddonsRequest { AddonsConfig: nullable.NewValue(types.GoogleCloudApigeeV1AddonsConfig { AdvancedApiOpsConfig: nullable.NewValue(types.GoogleCloudApigeeV1AdvancedApiOpsConfig { Enabled: nullable.NewValue(true) }), ApiSecurityConfig: nullable.NewValue(types.GoogleCloudApigeeV1ApiSecurityConfig { Enabled: nullable.NewValue(true), ExpiresAt: nullable.NewValue("string") }), ConnectorsPlatformConfig: nullable.NewValue(types.GoogleCloudApigeeV1ConnectorsPlatformConfig { Enabled: nullable.NewValue(true), ExpiresAt: nullable.NewValue("string") }), IntegrationConfig: nullable.NewValue(types.GoogleCloudApigeeV1IntegrationConfig { Enabled: nullable.NewValue(true) }), MonetizationConfig: nullable.NewValue(types.GoogleCloudApigeeV1MonetizationConfig { Enabled: nullable.NewValue(true) }) }) } })
```

    
### 
Creates an alias from a key/certificate pair. The structure of the request is controlled by the `format` query parameter: - `keycertfile` - Separate PEM-encoded key and certificate files are uploaded. Set `Content-Type: multipart/form-data` and include the `keyFile`, `certFile`, and `password` (if keys are encrypted) fields in the request body. If uploading to a truststore, omit `keyFile`. - `pkcs12` - A PKCS12 file is uploaded. Set `Content-Type: multipart/form-data`, provide the file in the `file` field, and include the `password` field if the file is encrypted in the request body. - `selfsignedcert` - A new private key and certificate are generated. Set `Content-Type: application/json` and include CertificateGenerationSpec in the request body.

**API Endpoint**: `POST /v1/{parent}/aliases`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	aliases "apigee_api/resources/v1/aliases"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Aliases.Create(aliases.CreateRequest { Parent: "string", Data: types.GoogleApiHttpBody { ContentType: nullable.NewValue("string"), Data: nullable.NewValue("string"), Extensions: nullable.NewValue([]map[string]interface{}{map[string]interface{}{}}) } })
```

    
### 
Create a Datastore for an org

**API Endpoint**: `POST /v1/{parent}/analytics/datastores`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	datastores "apigee_api/resources/v1/analytics/datastores"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Analytics.Datastores.Create(datastores.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1Datastore { CreateTime: nullable.NewValue("string"), DatastoreConfig: nullable.NewValue(types.GoogleCloudApigeeV1DatastoreConfig { BucketName: nullable.NewValue("string"), DatasetName: nullable.NewValue("string"), Path: nullable.NewValue("string"), ProjectId: nullable.NewValue("string"), TablePrefix: nullable.NewValue("string") }), DisplayName: nullable.NewValue("string"), LastUpdateTime: nullable.NewValue("string"), Org: nullable.NewValue("string"), Self: nullable.NewValue("string"), TargetType: nullable.NewValue("string") } })
```

    
### 
Test if Datastore configuration is correct. This includes checking if credentials provided by customer have required permissions in target destination storage

**API Endpoint**: `POST /v1/{parent}/analytics/datastores:test`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	datastores_test "apigee_api/resources/v1/analytics/datastores_test"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Analytics.DatastoresTest.Create(datastores_test.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1Datastore { CreateTime: nullable.NewValue("string"), DatastoreConfig: nullable.NewValue(types.GoogleCloudApigeeV1DatastoreConfig { BucketName: nullable.NewValue("string"), DatasetName: nullable.NewValue("string"), Path: nullable.NewValue("string"), ProjectId: nullable.NewValue("string"), TablePrefix: nullable.NewValue("string") }), DisplayName: nullable.NewValue("string"), LastUpdateTime: nullable.NewValue("string"), Org: nullable.NewValue("string"), Self: nullable.NewValue("string"), TargetType: nullable.NewValue("string") } })
```

    
### 
Submit a data export job to be processed in the background. If the request is successful, the API returns a 201 status, a URI that can be used to retrieve the status of the export job, and the `state` value of "enqueued".

**API Endpoint**: `POST /v1/{parent}/analytics/exports`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	exports "apigee_api/resources/v1/analytics/exports"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Analytics.Exports.Create(exports.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1ExportRequest { CsvDelimiter: nullable.NewValue("string"), DatastoreName: nullable.NewValue("string"), DateRange: nullable.NewValue(types.GoogleCloudApigeeV1DateRange { End: nullable.NewValue("string"), Start: nullable.NewValue("string") }), Description: nullable.NewValue("string"), Name: nullable.NewValue("string"), OutputFormat: nullable.NewValue("string") } })
```

    
### 
Creates a new category on the portal.

**API Endpoint**: `POST /v1/{parent}/apicategories`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	apicategories "apigee_api/resources/v1/apicategories"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Apicategories.Create(apicategories.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1ApiCategoryData { Id: nullable.NewValue("string"), Name: nullable.NewValue("string"), SiteId: nullable.NewValue("string"), UpdateTime: nullable.NewValue("string") } })
```

    
### 
Creates an API product in an organization. You create API products after you have proxied backend services using API proxies. An API product is a collection of API resources combined with quota settings and metadata that you can use to deliver customized and productized API bundles to your developer community. This metadata can include: - Scope - Environments - API proxies - Extensible profile API products enable you repackage APIs on the fly, without having to do any additional coding or configuration. Apigee recommends that you start with a simple API product including only required elements. You then provision credentials to apps to enable them to start testing your APIs. After you have authentication and authorization working against a simple API product, you can iterate to create finer-grained API products, defining different sets of API resources for each API product. **WARNING:** - If you don't specify an API proxy in the request body, *any* app associated with the product can make calls to *any* API in your entire organization. - If you don't specify an environment in the request body, the product allows access to all environments. For more information, see What is an API product?

**API Endpoint**: `POST /v1/{parent}/apiproducts`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	apiproducts "apigee_api/resources/v1/apiproducts"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Apiproducts.Create(apiproducts.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1ApiProduct { ApiResources: nullable.NewValue([]string{"string"}), ApprovalType: nullable.NewValue("string"), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), CreatedAt: nullable.NewValue("string"), Description: nullable.NewValue("string"), DisplayName: nullable.NewValue("string"), Environments: nullable.NewValue([]string{"string"}), GraphqlOperationGroup: nullable.NewValue(types.GoogleCloudApigeeV1GraphQlOperationGroup { OperationConfigType: nullable.NewValue("string"), OperationConfigs: nullable.NewValue([]types.GoogleCloudApigeeV1GraphQlOperationConfig{types.GoogleCloudApigeeV1GraphQlOperationConfig { ApiSource: nullable.NewValue("string"), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), Operations: nullable.NewValue([]types.GoogleCloudApigeeV1GraphQlOperation{types.GoogleCloudApigeeV1GraphQlOperation { Operation: nullable.NewValue("string"), OperationTypes: nullable.NewValue([]string{"string"}) }}), Quota: nullable.NewValue(types.GoogleCloudApigeeV1Quota { Interval: nullable.NewValue("string"), Limit: nullable.NewValue("string"), TimeUnit: nullable.NewValue("string") }) }}) }), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), OperationGroup: nullable.NewValue(types.GoogleCloudApigeeV1OperationGroup { OperationConfigType: nullable.NewValue("string"), OperationConfigs: nullable.NewValue([]types.GoogleCloudApigeeV1OperationConfig{types.GoogleCloudApigeeV1OperationConfig { ApiSource: nullable.NewValue("string"), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), Operations: nullable.NewValue([]types.GoogleCloudApigeeV1Operation{types.GoogleCloudApigeeV1Operation { Methods: nullable.NewValue([]string{"string"}), Resource: nullable.NewValue("string") }}), Quota: nullable.NewValue(types.GoogleCloudApigeeV1Quota { Interval: nullable.NewValue("string"), Limit: nullable.NewValue("string"), TimeUnit: nullable.NewValue("string") }) }}) }), Proxies: nullable.NewValue([]string{"string"}), Quota: nullable.NewValue("string"), QuotaCounterScope: nullable.NewValue(types.GoogleCloudApigeeV1ApiProductQuotaCounterScopeEnumOperation), QuotaInterval: nullable.NewValue("string"), QuotaTimeUnit: nullable.NewValue("string"), Scopes: nullable.NewValue([]string{"string"}) } })
```

    
### 
Creates an API proxy. The API proxy created will not be accessible at runtime until it is deployed to an environment. Create a new API proxy by setting the `name` query parameter to the name of the API proxy. Import an API proxy configuration bundle stored in zip format on your local machine to your organization by doing the following: * Set the `name` query parameter to the name of the API proxy. * Set the `action` query parameter to `import`. * Set the `Content-Type` header to `multipart/form-data`. * Pass as a file the name of API proxy configuration bundle stored in zip format on your local machine using the `file` form field. **Note**: To validate the API proxy configuration bundle only without importing it, set the `action` query parameter to `validate`. When importing an API proxy configuration bundle, if the API proxy does not exist, it will be created. If the API proxy exists, then a new revision is created. Invalid API proxy configurations are rejected, and a list of validation errors is returned to the client.

**API Endpoint**: `POST /v1/{parent}/apis`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	apis "apigee_api/resources/v1/apis"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Apis.Create(apis.CreateRequest { Parent: "string", Data: types.GoogleApiHttpBody { ContentType: nullable.NewValue("string"), Data: nullable.NewValue("string"), Extensions: nullable.NewValue([]map[string]interface{}{map[string]interface{}{}}) } })
```

    
### 
Creates an app associated with a developer. This API associates the developer app with the specified API product and auto-generates an API key for the app to use in calls to API proxies inside that API product. The `name` is the unique ID of the app that you can use in API calls. The `DisplayName` (set as an attribute) appears in the UI. If you don't set the `DisplayName` attribute, the `name` appears in the UI.

**API Endpoint**: `POST /v1/{parent}/apps`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	apps "apigee_api/resources/v1/apps"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Apps.Create(apps.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1DeveloperApp { ApiProducts: nullable.NewValue([]string{"string"}), AppFamily: nullable.NewValue("string"), AppId: nullable.NewValue("string"), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), CallbackUrl: nullable.NewValue("string"), CreatedAt: nullable.NewValue("string"), Credentials: nullable.NewValue([]types.GoogleCloudApigeeV1Credential{types.GoogleCloudApigeeV1Credential { ApiProducts: nullable.NewValue([]types.GoogleCloudApigeeV1ApiProductRef{types.GoogleCloudApigeeV1ApiProductRef { Apiproduct: nullable.NewValue("string"), Status: nullable.NewValue("string") }}), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), ConsumerKey: nullable.NewValue("string"), ConsumerSecret: nullable.NewValue("string"), ExpiresAt: nullable.NewValue("string"), IssuedAt: nullable.NewValue("string"), Scopes: nullable.NewValue([]string{"string"}), Status: nullable.NewValue("string") }}), DeveloperId: nullable.NewValue("string"), KeyExpiresIn: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), Scopes: nullable.NewValue([]string{"string"}), Status: nullable.NewValue("string") } })
```

    
### 
Creates a new ArchiveDeployment.

**API Endpoint**: `POST /v1/{parent}/archiveDeployments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	archive_deployments "apigee_api/resources/v1/archive_deployments"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.ArchiveDeployments.Create(archive_deployments.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1ArchiveDeployment { CreatedAt: nullable.NewValue("string"), GcsUri: nullable.NewValue("string"), Labels: nullable.NewValue(types.GoogleCloudApigeeV1ArchiveDeploymentLabels {  }), Name: nullable.NewValue("string"), Operation: nullable.NewValue("string"), UpdatedAt: nullable.NewValue("string") } })
```

    
### 
Generates a signed URL for uploading an Archive zip file to Google Cloud Storage. Once the upload is complete, the signed URL should be passed to CreateArchiveDeployment. When uploading to the generated signed URL, please follow these restrictions: * Source file type should be a zip file. * Source file size should not exceed 1GB limit. * No credentials should be attached - the signed URLs provide access to the target bucket using internal service identity; if credentials were attached, the identity from the credentials would be used, but that identity does not have permissions to upload files to the URL. When making a HTTP PUT request, these two headers need to be specified: * `content-type: application/zip` * `x-goog-content-length-range: 0,1073741824` And this header SHOULD NOT be specified: * `Authorization: Bearer YOUR_TOKEN`

**API Endpoint**: `POST /v1/{parent}/archiveDeployments:generateUploadUrl`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	archive_deployments_generate_upload_url "apigee_api/resources/v1/archive_deployments_generate_upload_url"
	url "net/url"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.ArchiveDeploymentsGenerateUploadUrl.Create(archive_deployments_generate_upload_url.CreateRequest { Parent: "string", Data: map[string]interface{}{} })
```

    
### 
Creates a new attachment of an environment to an instance. **Note:** Not supported for Apigee hybrid.

**API Endpoint**: `POST /v1/{parent}/attachments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	attachments "apigee_api/resources/v1/attachments"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Attachments.Create(attachments.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1InstanceAttachment { CreatedAt: nullable.NewValue("string"), Environment: nullable.NewValue("string"), Name: nullable.NewValue("string") } })
```

    
### 
Updates developer attributes. This API replaces the existing attributes with those specified in the request. Add new attributes, and include or exclude any existing attributes that you want to retain or remove, respectively. The custom attribute limit is 18. **Note**: OAuth access tokens and Key Management Service (KMS) entities (apps, developers, and API products) are cached for 180 seconds (default). Any custom attributes associated with these entities are cached for at least 180 seconds after the entity is accessed at runtime. Therefore, an `ExpiresIn` element on the OAuthV2 policy won't be able to expire an access token in less than 180 seconds.

**API Endpoint**: `POST /v1/{parent}/attributes`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	attributes "apigee_api/resources/v1/attributes"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Attributes.Create0(attributes.Create0Request { Parent: "string", Data: types.GoogleCloudApigeeV1Attributes { Attribute: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}) } })
```

    
### 
Creates a new canary evaluation for an organization.

**API Endpoint**: `POST /v1/{parent}/canaryevaluations`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	canaryevaluations "apigee_api/resources/v1/canaryevaluations"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Canaryevaluations.Create(canaryevaluations.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1CanaryEvaluation { Control: nullable.NewValue("string"), CreateTime: nullable.NewValue("string"), EndTime: nullable.NewValue("string"), MetricLabels: nullable.NewValue(types.GoogleCloudApigeeV1CanaryEvaluationMetricLabels { Env: nullable.NewValue("string"), InstanceId: nullable.NewValue("string"), Location: nullable.NewValue("string") }), Name: nullable.NewValue("string"), StartTime: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1CanaryEvaluationStateEnumRunning), Treatment: nullable.NewValue("string"), Verdict: nullable.NewValue(types.GoogleCloudApigeeV1CanaryEvaluationVerdictEnumFail) } })
```

    
### 
Creates a new data collector.

**API Endpoint**: `POST /v1/{parent}/datacollectors`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	datacollectors "apigee_api/resources/v1/datacollectors"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Datacollectors.Create(datacollectors.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1DataCollector { CreatedAt: nullable.NewValue("string"), Description: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), Type: nullable.NewValue(types.GoogleCloudApigeeV1DataCollectorTypeEnumBoolean) } })
```

    
### 
Creates a debug session for a deployed API Proxy revision.

**API Endpoint**: `POST /v1/{parent}/debugsessions`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	debugsessions "apigee_api/resources/v1/debugsessions"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Debugsessions.Create(debugsessions.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1DebugSession { Count: nullable.NewValue(123), CreateTime: nullable.NewValue("string"), Filter: nullable.NewValue("string"), Name: nullable.NewValue("string"), Timeout: nullable.NewValue("string"), Tracesize: nullable.NewValue(123), Validity: nullable.NewValue(123) } })
```

    
### 
Creates a developer. Once created, the developer can register an app and obtain an API key. At creation time, a developer is set as `active`. To change the developer status, use the SetDeveloperStatus API.

**API Endpoint**: `POST /v1/{parent}/developers`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	developers "apigee_api/resources/v1/developers"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Developers.Create(developers.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1Developer { AccessType: nullable.NewValue("string"), AppFamily: nullable.NewValue("string"), Apps: nullable.NewValue([]string{"string"}), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), Companies: nullable.NewValue([]string{"string"}), CreatedAt: nullable.NewValue("string"), DeveloperId: nullable.NewValue("string"), Email: nullable.NewValue("string"), FirstName: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), LastName: nullable.NewValue("string"), OrganizationName: nullable.NewValue("string"), Status: nullable.NewValue("string"), UserName: nullable.NewValue("string") } })
```

    
### 
Creates an endpoint attachment. **Note:** Not supported for Apigee hybrid.

**API Endpoint**: `POST /v1/{parent}/endpointAttachments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	endpoint_attachments "apigee_api/resources/v1/endpoint_attachments"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.EndpointAttachments.Create(endpoint_attachments.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1EndpointAttachment { ConnectionState: nullable.NewValue(types.GoogleCloudApigeeV1EndpointAttachmentConnectionStateEnumAccepted), Host: nullable.NewValue("string"), Location: nullable.NewValue("string"), Name: nullable.NewValue("string"), ServiceAttachment: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1EndpointAttachmentStateEnumActive) } })
```

    
### 
Creates key value entries in a key value map scoped to an organization, environment, or API proxy. **Note**: Supported for Apigee hybrid 1.8.x and higher.

**API Endpoint**: `POST /v1/{parent}/entries`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	entries "apigee_api/resources/v1/entries"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Entries.Create(entries.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1KeyValueEntry { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") } })
```

    
### 
Creates a new environment group.

**API Endpoint**: `POST /v1/{parent}/envgroups`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	envgroups "apigee_api/resources/v1/envgroups"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Envgroups.Create(envgroups.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1EnvironmentGroup { CreatedAt: nullable.NewValue("string"), Hostnames: nullable.NewValue([]string{"string"}), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1EnvironmentGroupStateEnumActive) } })
```

    
### 
CreateSecurityProfileEnvironmentAssociation creates profile environment association i.e. attaches environment to security profile.

**API Endpoint**: `POST /v1/{parent}/environments`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	environments "apigee_api/resources/v1/environments"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Environments.Create(environments.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1SecurityProfileEnvironmentAssociation { AttachTime: nullable.NewValue("string"), Name: nullable.NewValue("string"), SecurityProfileRevisionId: nullable.NewValue("string") } })
```

    
### 
Submit a query at host level to be processed in the background. If the submission of the query succeeds, the API returns a 201 status and an ID that refer to the query. In addition to the HTTP status 201, the `state` of "enqueued" means that the request succeeded.

**API Endpoint**: `POST /v1/{parent}/hostQueries`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	host_queries "apigee_api/resources/v1/host_queries"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.HostQueries.Create(host_queries.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1Query { CsvDelimiter: nullable.NewValue("string"), Dimensions: nullable.NewValue([]string{"string"}), EnvgroupHostname: nullable.NewValue("string"), Filter: nullable.NewValue("string"), GroupByTimeUnit: nullable.NewValue("string"), Limit: nullable.NewValue(123), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1QueryMetric{types.GoogleCloudApigeeV1QueryMetric { Alias: nullable.NewValue("string"), Function: nullable.NewValue("string"), Name: nullable.NewValue("string"), Operator: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), Name: nullable.NewValue("string"), OutputFormat: nullable.NewValue("string"), ReportDefinitionId: nullable.NewValue("string"), TimeRange: nullable.NewValue("could be anything") } })
```

    
### 
Submit a query at host level to be processed in the background. If the submission of the query succeeds, the API returns a 201 status and an ID that refer to the query. In addition to the HTTP status 201, the `state` of "enqueued" means that the request succeeded.

**API Endpoint**: `POST /v1/{parent}/hostSecurityReports`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	host_security_reports "apigee_api/resources/v1/host_security_reports"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.HostSecurityReports.Create(host_security_reports.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1SecurityReportQuery { CsvDelimiter: nullable.NewValue("string"), Dimensions: nullable.NewValue([]string{"string"}), DisplayName: nullable.NewValue("string"), EnvgroupHostname: nullable.NewValue("string"), Filter: nullable.NewValue("string"), GroupByTimeUnit: nullable.NewValue("string"), Limit: nullable.NewValue(123), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1SecurityReportQueryMetric{types.GoogleCloudApigeeV1SecurityReportQueryMetric { AggregationFunction: nullable.NewValue("string"), Alias: nullable.NewValue("string"), Name: nullable.NewValue("string"), Operator: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), MimeType: nullable.NewValue("string"), ReportDefinitionId: nullable.NewValue("string"), TimeRange: nullable.NewValue("could be anything") } })
```

    
### 
Creates an Apigee runtime instance. The instance is accessible from the authorized network configured on the organization. **Note:** Not supported for Apigee hybrid.

**API Endpoint**: `POST /v1/{parent}/instances`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	instances "apigee_api/resources/v1/instances"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Instances.Create(instances.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1Instance { ConsumerAcceptList: nullable.NewValue([]string{"string"}), CreatedAt: nullable.NewValue("string"), Description: nullable.NewValue("string"), DiskEncryptionKeyName: nullable.NewValue("string"), DisplayName: nullable.NewValue("string"), Host: nullable.NewValue("string"), IpRange: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), Location: nullable.NewValue("string"), Name: nullable.NewValue("string"), PeeringCidrRange: nullable.NewValue(types.GoogleCloudApigeeV1InstancePeeringCidrRangeEnumCidrRangeUnspecified), Port: nullable.NewValue("string"), RuntimeVersion: nullable.NewValue("string"), ServiceAttachment: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1InstanceStateEnumActive) } })
```

    
### 
Creates a custom consumer key and secret for a developer app. This is particularly useful if you want to migrate existing consumer keys and secrets to Apigee from another system. Consumer keys and secrets can contain letters, numbers, underscores, and hyphens. No other special characters are allowed. To avoid service disruptions, a consumer key and secret should not exceed 2 KBs each. **Note**: When creating the consumer key and secret, an association to API products will not be made. Therefore, you should not specify the associated API products in your request. Instead, use the UpdateDeveloperAppKey API to make the association after the consumer key and secret are created. If a consumer key and secret already exist, you can keep them or delete them using the DeleteDeveloperAppKey API.

**API Endpoint**: `POST /v1/{parent}/keys`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	keys "apigee_api/resources/v1/keys"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Keys.Create(keys.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1DeveloperAppKey { ApiProducts: nullable.NewValue([]interface{}{"could be anything"}), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), ConsumerKey: nullable.NewValue("string"), ConsumerSecret: nullable.NewValue("string"), ExpiresAt: nullable.NewValue("string"), ExpiresInSeconds: nullable.NewValue("string"), IssuedAt: nullable.NewValue("string"), Scopes: nullable.NewValue([]string{"string"}), Status: nullable.NewValue("string") } })
```

    
### 
Creates a custom consumer key and secret for a developer app. This is particularly useful if you want to migrate existing consumer keys and secrets to Apigee from another system. Consumer keys and secrets can contain letters, numbers, underscores, and hyphens. No other special characters are allowed. To avoid service disruptions, a consumer key and secret should not exceed 2 KBs each. **Note**: When creating the consumer key and secret, an association to API products will not be made. Therefore, you should not specify the associated API products in your request. Instead, use the UpdateDeveloperAppKey API to make the association after the consumer key and secret are created. If a consumer key and secret already exist, you can keep them or delete them using the DeleteDeveloperAppKey API.

**API Endpoint**: `POST /v1/{parent}/keys/create`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	create "apigee_api/resources/v1/keys/create"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Keys.Create.Create(create.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1DeveloperAppKey { ApiProducts: nullable.NewValue([]interface{}{"could be anything"}), Attributes: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), ConsumerKey: nullable.NewValue("string"), ConsumerSecret: nullable.NewValue("string"), ExpiresAt: nullable.NewValue("string"), ExpiresInSeconds: nullable.NewValue("string"), IssuedAt: nullable.NewValue("string"), Scopes: nullable.NewValue([]string{"string"}), Status: nullable.NewValue("string") } })
```

    
### 
Creates a keystore or truststore. - Keystore: Contains certificates and their associated keys. - Truststore: Contains trusted certificates used to validate a server's certificate. These certificates are typically self-signed certificates or certificates that are not signed by a trusted CA.

**API Endpoint**: `POST /v1/{parent}/keystores`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	keystores "apigee_api/resources/v1/keystores"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Keystores.Create(keystores.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1Keystore { Aliases: nullable.NewValue([]string{"string"}), Name: nullable.NewValue("string") } })
```

    
### 
Creates a key value map in an organization.

**API Endpoint**: `POST /v1/{parent}/keyvaluemaps`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	keyvaluemaps "apigee_api/resources/v1/keyvaluemaps"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Keyvaluemaps.Create(keyvaluemaps.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1KeyValueMap { Encrypted: nullable.NewValue(true), Name: nullable.NewValue("string") } })
```

    
### 
Creates a NAT address. The address is created in the RESERVED state and a static external IP address will be provisioned. At this time, the instance will not use this IP address for Internet egress traffic. The address can be activated for use once any required firewall IP whitelisting has been completed. **Note:** Not supported for Apigee hybrid.

**API Endpoint**: `POST /v1/{parent}/natAddresses`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	nat_addresses "apigee_api/resources/v1/nat_addresses"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.NatAddresses.Create(nat_addresses.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1NatAddress { IpAddress: nullable.NewValue("string"), Name: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1NatAddressStateEnumActive) } })
```

    
### 
Creates a trace configuration override. The response contains a system-generated UUID, that can be used to view, update, or delete the configuration override. Use the List API to view the existing trace configuration overrides.

**API Endpoint**: `POST /v1/{parent}/overrides`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	overrides "apigee_api/resources/v1/overrides"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Overrides.Create(overrides.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1TraceConfigOverride { ApiProxy: nullable.NewValue("string"), Name: nullable.NewValue("string"), SamplingConfig: nullable.NewValue(types.GoogleCloudApigeeV1TraceSamplingConfig { Sampler: nullable.NewValue(types.GoogleCloudApigeeV1TraceSamplingConfigSamplerEnumOff), SamplingRate: nullable.NewValue(123.45) }) } })
```

    
### 
Submit a query to be processed in the background. If the submission of the query succeeds, the API returns a 201 status and an ID that refer to the query. In addition to the HTTP status 201, the `state` of "enqueued" means that the request succeeded.

**API Endpoint**: `POST /v1/{parent}/queries`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	queries "apigee_api/resources/v1/queries"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Queries.Create(queries.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1Query { CsvDelimiter: nullable.NewValue("string"), Dimensions: nullable.NewValue([]string{"string"}), EnvgroupHostname: nullable.NewValue("string"), Filter: nullable.NewValue("string"), GroupByTimeUnit: nullable.NewValue("string"), Limit: nullable.NewValue(123), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1QueryMetric{types.GoogleCloudApigeeV1QueryMetric { Alias: nullable.NewValue("string"), Function: nullable.NewValue("string"), Name: nullable.NewValue("string"), Operator: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), Name: nullable.NewValue("string"), OutputFormat: nullable.NewValue("string"), ReportDefinitionId: nullable.NewValue("string"), TimeRange: nullable.NewValue("could be anything") } })
```

    
### 
Create a rate plan that is associated with an API product in an organization. Using rate plans, API product owners can monetize their API products by configuring one or more of the following: - Billing frequency - Initial setup fees for using an API product - Payment funding model (postpaid only) - Fixed recurring or consumption-based charges for using an API product - Revenue sharing with developer partners An API product can have multiple rate plans associated with it but *only one* rate plan can be active at any point of time. **Note: From the developer's perspective, they purchase API products not rate plans.

**API Endpoint**: `POST /v1/{parent}/rateplans`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	rateplans "apigee_api/resources/v1/rateplans"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Rateplans.Create(rateplans.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1RatePlan { Apiproduct: nullable.NewValue("string"), BillingPeriod: nullable.NewValue(types.GoogleCloudApigeeV1RatePlanBillingPeriodEnumBillingPeriodUnspecified), ConsumptionPricingRates: nullable.NewValue([]types.GoogleCloudApigeeV1RateRange{types.GoogleCloudApigeeV1RateRange { End: nullable.NewValue("string"), Fee: nullable.NewValue(types.GoogleTypeMoney { CurrencyCode: nullable.NewValue("string"), Nanos: nullable.NewValue(123), Units: nullable.NewValue("string") }), Start: nullable.NewValue("string") }}), ConsumptionPricingType: nullable.NewValue(types.GoogleCloudApigeeV1RatePlanConsumptionPricingTypeEnumBanded), CreatedAt: nullable.NewValue("string"), CurrencyCode: nullable.NewValue("string"), Description: nullable.NewValue("string"), DisplayName: nullable.NewValue("string"), EndTime: nullable.NewValue("string"), FixedFeeFrequency: nullable.NewValue(123), FixedRecurringFee: nullable.NewValue(types.GoogleTypeMoney { CurrencyCode: nullable.NewValue("string"), Nanos: nullable.NewValue(123), Units: nullable.NewValue("string") }), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), PaymentFundingModel: nullable.NewValue(types.GoogleCloudApigeeV1RatePlanPaymentFundingModelEnumPaymentFundingModelUnspecified), RevenueShareRates: nullable.NewValue([]types.GoogleCloudApigeeV1RevenueShareRange{types.GoogleCloudApigeeV1RevenueShareRange { End: nullable.NewValue("string"), SharePercentage: nullable.NewValue(123.45), Start: nullable.NewValue("string") }}), RevenueShareType: nullable.NewValue(types.GoogleCloudApigeeV1RatePlanRevenueShareTypeEnumFixed), SetupFee: nullable.NewValue(types.GoogleTypeMoney { CurrencyCode: nullable.NewValue("string"), Nanos: nullable.NewValue(123), Units: nullable.NewValue("string") }), StartTime: nullable.NewValue("string"), State: nullable.NewValue(types.GoogleCloudApigeeV1RatePlanStateEnumDraft) } })
```

    
### 
Creates a Reference in the specified environment.

**API Endpoint**: `POST /v1/{parent}/references`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	references "apigee_api/resources/v1/references"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.References.Create(references.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1Reference { Description: nullable.NewValue("string"), Name: nullable.NewValue("string"), Refers: nullable.NewValue("string"), ResourceType: nullable.NewValue("string") } })
```

    
### 
Creates a Custom Report for an Organization. A Custom Report provides Apigee Customers to create custom dashboards in addition to the standard dashboards which are provided. The Custom Report in its simplest form contains specifications about metrics, dimensions and filters. It is important to note that the custom report by itself does not provide an executable entity. The Edge UI converts the custom report definition into an analytics query and displays the result in a chart.

**API Endpoint**: `POST /v1/{parent}/reports`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	reports "apigee_api/resources/v1/reports"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Reports.Create(reports.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1CustomReport { ChartType: nullable.NewValue("string"), Comments: nullable.NewValue([]string{"string"}), CreatedAt: nullable.NewValue("string"), Dimensions: nullable.NewValue([]string{"string"}), DisplayName: nullable.NewValue("string"), Environment: nullable.NewValue("string"), Filter: nullable.NewValue("string"), FromTime: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), LastViewedAt: nullable.NewValue("string"), Limit: nullable.NewValue("string"), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1CustomReportMetric{types.GoogleCloudApigeeV1CustomReportMetric { Function: nullable.NewValue("string"), Name: nullable.NewValue("string") }}), Name: nullable.NewValue("string"), Offset: nullable.NewValue("string"), Organization: nullable.NewValue("string"), Properties: nullable.NewValue([]types.GoogleCloudApigeeV1ReportProperty{types.GoogleCloudApigeeV1ReportProperty { Property: nullable.NewValue("string"), Value: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}) }}), SortByCols: nullable.NewValue([]string{"string"}), SortOrder: nullable.NewValue("string"), Tags: nullable.NewValue([]string{"string"}), TimeUnit: nullable.NewValue("string"), ToTime: nullable.NewValue("string"), Topk: nullable.NewValue("string") } })
```

    
### 
Creates a resource file. Specify the `Content-Type` as `application/octet-stream` or `multipart/form-data`. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).

**API Endpoint**: `POST /v1/{parent}/resourcefiles`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	resourcefiles "apigee_api/resources/v1/resourcefiles"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Resourcefiles.Create(resourcefiles.CreateRequest { Parent: "string", Data: types.GoogleApiHttpBody { ContentType: nullable.NewValue("string"), Data: nullable.NewValue("string"), Extensions: nullable.NewValue([]map[string]interface{}{map[string]interface{}{}}) } })
```

    
### 
Submit a report request to be processed in the background. If the submission succeeds, the API returns a 200 status and an ID that refer to the report request. In addition to the HTTP status 200, the `state` of "enqueued" means that the request succeeded.

**API Endpoint**: `POST /v1/{parent}/securityReports`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	security_reports "apigee_api/resources/v1/security_reports"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.SecurityReports.Create(security_reports.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1SecurityReportQuery { CsvDelimiter: nullable.NewValue("string"), Dimensions: nullable.NewValue([]string{"string"}), DisplayName: nullable.NewValue("string"), EnvgroupHostname: nullable.NewValue("string"), Filter: nullable.NewValue("string"), GroupByTimeUnit: nullable.NewValue("string"), Limit: nullable.NewValue(123), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1SecurityReportQueryMetric{types.GoogleCloudApigeeV1SecurityReportQueryMetric { AggregationFunction: nullable.NewValue("string"), Alias: nullable.NewValue("string"), Name: nullable.NewValue("string"), Operator: nullable.NewValue("string"), Value: nullable.NewValue("string") }}), MimeType: nullable.NewValue("string"), ReportDefinitionId: nullable.NewValue("string"), TimeRange: nullable.NewValue("could be anything") } })
```

    
### 
Uploads a ZIP-formatted shared flow configuration bundle to an organization. If the shared flow already exists, this creates a new revision of it. If the shared flow does not exist, this creates it. Once imported, the shared flow revision must be deployed before it can be accessed at runtime. The size limit of a shared flow bundle is 15 MB.

**API Endpoint**: `POST /v1/{parent}/sharedflows`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	sharedflows "apigee_api/resources/v1/sharedflows"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Sharedflows.Create(sharedflows.CreateRequest { Parent: "string", Data: types.GoogleApiHttpBody { ContentType: nullable.NewValue("string"), Data: nullable.NewValue("string"), Extensions: nullable.NewValue([]map[string]interface{}{map[string]interface{}{}}) } })
```

    
### 
Creates a subscription to an API product. 

**API Endpoint**: `POST /v1/{parent}/subscriptions`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	subscriptions "apigee_api/resources/v1/subscriptions"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Subscriptions.Create(subscriptions.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1DeveloperSubscription { Apiproduct: nullable.NewValue("string"), CreatedAt: nullable.NewValue("string"), EndTime: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), Name: nullable.NewValue("string"), StartTime: nullable.NewValue("string") } })
```

    
### 
Creates a TargetServer in the specified environment.

**API Endpoint**: `POST /v1/{parent}/targetservers`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	targetservers "apigee_api/resources/v1/targetservers"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Targetservers.Create(targetservers.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1TargetServer { Description: nullable.NewValue("string"), Host: nullable.NewValue("string"), IsEnabled: nullable.NewValue(true), Name: nullable.NewValue("string"), Port: nullable.NewValue(123), Protocol: nullable.NewValue(types.GoogleCloudApigeeV1TargetServerProtocolEnumGrpc), SSlInfo: nullable.NewValue(types.GoogleCloudApigeeV1TlsInfo { Ciphers: nullable.NewValue([]string{"string"}), ClientAuthEnabled: nullable.NewValue(true), CommonName: nullable.NewValue(types.GoogleCloudApigeeV1TlsInfoCommonName { Value: nullable.NewValue("string"), WildcardMatch: nullable.NewValue(true) }), Enabled: nullable.NewValue(true), IgnoreValidationErrors: nullable.NewValue(true), KeyAlias: nullable.NewValue("string"), KeyStore: nullable.NewValue("string"), Protocols: nullable.NewValue([]string{"string"}), TrustStore: nullable.NewValue("string") }) } })
```

    
### 
Creates a subscription for the environment's Pub/Sub topic. The server will assign a random name for this subscription. The "name" and "push_config" must *not* be specified.

**API Endpoint**: `POST /v1/{parent}:subscribe`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	subscribe "apigee_api/resources/v1/subscribe"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Subscribe.Create(subscribe.CreateRequest { Parent: "string" })
```

    
### 
Deletes a subscription for the environment's Pub/Sub topic.

**API Endpoint**: `POST /v1/{parent}:unsubscribe`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	unsubscribe "apigee_api/resources/v1/unsubscribe"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Unsubscribe.Create(unsubscribe.CreateRequest { Parent: "string", Data: types.GoogleCloudApigeeV1Subscription { Name: nullable.NewValue("string") } })
```

    
### 
ComputeEnvironmentScores calculates scores for requested time range for the specified security profile and environment.

**API Endpoint**: `POST /v1/{profileEnvironment}:computeEnvironmentScores`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	compute_environment_scores "apigee_api/resources/v1/compute_environment_scores"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.ComputeEnvironmentScores.Create(compute_environment_scores.CreateRequest { ProfileEnvironment: "string", Data: types.GoogleCloudApigeeV1ComputeEnvironmentScoresRequest { Filters: nullable.NewValue([]types.GoogleCloudApigeeV1ComputeEnvironmentScoresRequestFilter{types.GoogleCloudApigeeV1ComputeEnvironmentScoresRequestFilter { ScorePath: nullable.NewValue("string") }}), PageSize: nullable.NewValue(123), PageToken: nullable.NewValue("string"), TimeRange: nullable.NewValue(types.GoogleTypeInterval { EndTime: nullable.NewValue("string"), StartTime: nullable.NewValue("string") }) } })
```

    
### 
Provisions a new Apigee organization with a functioning runtime. This is the standard way to create trial organizations for a free Apigee trial.

**API Endpoint**: `POST /v1/{project}:provisionOrganization`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	provision_organization "apigee_api/resources/v1/provision_organization"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.ProvisionOrganization.Create(provision_organization.CreateRequest { Project: "string", Data: types.GoogleCloudApigeeV1ProvisionOrganizationRequest { AnalyticsRegion: nullable.NewValue("string"), AuthorizedNetwork: nullable.NewValue("string"), RuntimeLocation: nullable.NewValue("string") } })
```

    
### 
Sets the IAM policy on an environment, if the policy already exists it will be replaced. For more information, see [Manage users, roles, and permissions using the API](https://cloud.google.com/apigee/docs/api-platform/system-administration/manage-users-roles). You must have the `apigee.environments.setIamPolicy` permission to call this API.

**API Endpoint**: `POST /v1/{resource}:setIamPolicy`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	set_iam_policy "apigee_api/resources/v1/set_iam_policy"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.SetIamPolicy.Create(set_iam_policy.CreateRequest { Resource: "string", Data: types.GoogleIamV1SetIamPolicyRequest { Policy: nullable.NewValue(types.GoogleIamV1Policy { AuditConfigs: nullable.NewValue([]types.GoogleIamV1AuditConfig{types.GoogleIamV1AuditConfig { AuditLogConfigs: nullable.NewValue([]types.GoogleIamV1AuditLogConfig{types.GoogleIamV1AuditLogConfig { ExemptedMembers: nullable.NewValue([]string{"string"}), LogType: nullable.NewValue(types.GoogleIamV1AuditLogConfigLogTypeEnumAdminRead) }}), Service: nullable.NewValue("string") }}), Bindings: nullable.NewValue([]types.GoogleIamV1Binding{types.GoogleIamV1Binding { Condition: nullable.NewValue(types.GoogleTypeExpr { Description: nullable.NewValue("string"), Expression: nullable.NewValue("string"), Location: nullable.NewValue("string"), Title: nullable.NewValue("string") }), Members: nullable.NewValue([]string{"string"}), Role: nullable.NewValue("string") }}), Etag: nullable.NewValue("string"), Version: nullable.NewValue(123) }), UpdateMask: nullable.NewValue("string") } })
```

    
### 
Tests the permissions of a user on an environment, and returns a subset of permissions that the user has on the environment. If the environment does not exist, an empty permission set is returned (a NOT_FOUND error is not returned).

**API Endpoint**: `POST /v1/{resource}:testIamPermissions`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	test_iam_permissions "apigee_api/resources/v1/test_iam_permissions"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.TestIamPermissions.Create(test_iam_permissions.CreateRequest { Resource: "string", Data: types.GoogleIamV1TestIamPermissionsRequest { Permissions: nullable.NewValue([]string{"string"}) } })
```

    
### 
Update an existing custom report definition

**API Endpoint**: `PUT /v1/{name}`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	v1 "apigee_api/resources/v1"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Put(v1.PutRequest { Name: "string", Data: types.GoogleCloudApigeeV1CustomReport { ChartType: nullable.NewValue("string"), Comments: nullable.NewValue([]string{"string"}), CreatedAt: nullable.NewValue("string"), Dimensions: nullable.NewValue([]string{"string"}), DisplayName: nullable.NewValue("string"), Environment: nullable.NewValue("string"), Filter: nullable.NewValue("string"), FromTime: nullable.NewValue("string"), LastModifiedAt: nullable.NewValue("string"), LastViewedAt: nullable.NewValue("string"), Limit: nullable.NewValue("string"), Metrics: nullable.NewValue([]types.GoogleCloudApigeeV1CustomReportMetric{types.GoogleCloudApigeeV1CustomReportMetric { Function: nullable.NewValue("string"), Name: nullable.NewValue("string") }}), Name: nullable.NewValue("string"), Offset: nullable.NewValue("string"), Organization: nullable.NewValue("string"), Properties: nullable.NewValue([]types.GoogleCloudApigeeV1ReportProperty{types.GoogleCloudApigeeV1ReportProperty { Property: nullable.NewValue("string"), Value: nullable.NewValue([]types.GoogleCloudApigeeV1Attribute{types.GoogleCloudApigeeV1Attribute { Name: nullable.NewValue("string"), Value: nullable.NewValue("string") }}) }}), SortByCols: nullable.NewValue([]string{"string"}), SortOrder: nullable.NewValue("string"), Tags: nullable.NewValue([]string{"string"}), TimeUnit: nullable.NewValue("string"), ToTime: nullable.NewValue("string"), Topk: nullable.NewValue("string") } })
```

    
### 
Updates a resource file. Specify the `Content-Type` as `application/octet-stream` or `multipart/form-data`. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).

**API Endpoint**: `PUT /v1/{parent}/resourcefiles/{type}/{name}`


#### Example Snippet

```go
import (
	sdk "apigee_api/client"
	os "os"
	x "apigee_api/resources/v1/resourcefiles/x"
	types "apigee_api/types"
	nullable "apigee_api/nullable"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))

res, err := client.V1.Resourcefiles.X.Put(x.PutRequest { Parent: "string", Type: "string", Name: "string", Data: types.GoogleApiHttpBody { ContentType: nullable.NewValue("string"), Data: nullable.NewValue("string"), Extensions: nullable.NewValue([]map[string]interface{}{map[string]interface{}{}}) } })
```

    