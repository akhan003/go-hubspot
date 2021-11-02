# Go API client for owners

HubSpot uses **owners** to assign CRM objects to specific people in your organization. The endpoints described here are used to get a list of the owners that are available for an account. To assign an owner to an object, set the hubspot_owner_id property using the appropriate CRM object update or create a request.  If teams are available for your HubSpot tier, these endpoints will also indicate which team an owner belongs to. Team membership can be one of PRIMARY (default), SECONDARY, or CHILD.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v3
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./owners"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hubapi.com/*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*OwnersApi* | [**Getcrmv3ownersGetPage**](docs/OwnersApi.md#getcrmv3ownersgetpage) | **Get** /crm/v3/owners/ | Get a page of owners
*OwnersApi* | [**Getcrmv3ownersownerIdGetById**](docs/OwnersApi.md#getcrmv3ownersowneridgetbyid) | **Get** /crm/v3/owners/{ownerId} | Read an owner by given &#x60;id&#x60; or &#x60;userId&#x60;

## Documentation For Models

 - [CollectionResponsePublicOwnerForwardPaging](docs/CollectionResponsePublicOwnerForwardPaging.md)
 - [ErrorDetail](docs/ErrorDetail.md)
 - [ForwardPaging](docs/ForwardPaging.md)
 - [ModelError](docs/ModelError.md)
 - [NextPage](docs/NextPage.md)
 - [PublicOwner](docs/PublicOwner.md)
 - [PublicTeam](docs/PublicTeam.md)

## Documentation For Authorization

## hapikey
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```
## oauth2
- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **crm.objects.owners.read**:  

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.
```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```
## oauth2_legacy
- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.hubspot.com/oauth/authorize
- **Scopes**: 
 - **contacts**: Read from and write to my Contacts

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.
```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```

## Author

