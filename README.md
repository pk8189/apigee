
# Apigee API Go SDK


## Overview
Use the Apigee API to programmatically develop and manage APIs with a set of RESTful operations. Develop and secure API proxies, deploy and undeploy API proxy revisions, monitor APIs, configure environments, manage users, and more. Note: This product is available as a free trial for a time period of 60 days.


### Example Client Initialization

```go
import (
	sdk "apigee_api/client"
	os "os"
)

client := sdk.NewClient(sdk.WithOauth2(os.Getenv("API_TOKEN")), sdk.WithOauth2c(os.Getenv("API_TOKEN")))
```

### SDK Usage 
 See [SDK Examples](SDK_EXAMPLES.md) for example usage of all SDK functionality