# go SDK for amazon-sp-api

Port of [https://github.com/amz-tools/amazon-sp-api](https://github.com/amz-tools/amazon-sp-api).
Credits due to the contributors of that repository: [https://github.com/amz-tools](https://github.com/amz-tools)

The client handles calls to the Amazon Selling Partner API. It wraps up all the necessary stuff 
such as requesting access token, security token and signing requests with AWS4 signature.

## Prerequisites
Make sure that you followed the [Selling Partner API Developer Guide](https://github.com/amzn/selling-partner-api-docs/blob/main/guides/developer-guide/SellingPartnerApiDeveloperGuide.md)
and have successfully completed the steps [Registering as a developer](https://github.com/amzn/selling-partner-api-docs/blob/main/guides/developer-guide/SellingPartnerApiDeveloperGuide.md#registering-as-a-developer),
[Registering your Selling Partner API application](https://github.com/amzn/selling-partner-api-docs/blob/main/guides/developer-guide/SellingPartnerApiDeveloperGuide.md#registering-your-selling-partner-api-application)
and have a valid refresh_token (if you use the client only for your own seller account the easiest way is using the
self authorization as described in the developer guide).

## Installation
```bash
go get -u github.com/amazinsellers/amazon-sp-api-sdk-go
```

## Getting Started
Before you can use the client you need to add your app client and aws user credentials as environment variables:

* `SELLING_PARTNER_APP_CLIENT_ID`=<YOUR_APP_CLIENT_ID> ([see SP Developer Guide "Viewing your developer information"](https://github.com/amzn/selling-partner-api-docs/blob/main/guides/developer-guide/SellingPartnerApiDeveloperGuide.md#viewing-your-developer-information))
* `SELLING_PARTNER_APP_CLIENT_SECRET`=<YOUR_APP_CLIENT_SECRET> ([see SP Developer Guide "Viewing your developer information"](https://github.com/amzn/selling-partner-api-docs/blob/main/guides/developer-guide/SellingPartnerApiDeveloperGuide.md#viewing-your-developer-information))
* `AWS_SELLING_PARTNER_ACCESS_KEY_ID` or `AWS_ACCESS_KEY_ID`=<YOUR_AWS_USER_ID> ([see SP Developer Guide "Create an IAM user"](https://github.com/amzn/selling-partner-api-docs/blob/main/guides/developer-guide/SellingPartnerApiDeveloperGuide.md#step-2-create-an-iam-user))
* `AWS_SELLING_PARTNER_SECRET_ACCESS_KEY` or `AWS_SECRET_ACCESS_KEY`=<YOUR_AWS_USER_SECRET> ([see SP Developer Guide "Create an IAM user"](https://github.com/amzn/selling-partner-api-docs/blob/main/guides/developer-guide/SellingPartnerApiDeveloperGuide.md#step-2-create-an-iam-user))
* `AWS_SELLING_PARTNER_ROLE`=<YOUR_AWS_SELLING_PARTNER_API_ROLE> ([see SP Developer Guide "Create an IAM role"](https://github.com/amzn/selling-partner-api-docs/blob/main/guides/developer-guide/SellingPartnerApiDeveloperGuide.md#step-4-create-an-iam-role))

