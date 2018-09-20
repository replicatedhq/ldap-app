# Replicated LDAP Example App

This Replicated app demonstrates how LDAP sync functionality can be implemented with [Provisioning API](https://help.replicated.com/api/integration-api/provisioning-api/).

## Building

Ensure that GOPATH is set.

```bash
GOOS=linux go build .
docker build -t registry.replicated.com/<slug>/ldap-app .
docker push registry.replicated.com/<slug>/ldap-app
```

`<slug>` is unique for every Replicated app.  It can be found on the `Images` page in [Vendor Web](https://vendor.replicated.com) for each app.

## Creating the app

 - Create a new app in [Vendor Web](https://vendor.replicated.com).
 - Create a new release and copy the contents of the [replicated.yaml](./replicated.yaml) file.
 - Promote the release to a channel.
 - Create and download a license assigned to this channel
 - Install Replicated following [these instructions](https://help.replicated.com/docs/native/customer-installations/installing/).
 - Once the license is installed, configure LDAP settings and start the app.
