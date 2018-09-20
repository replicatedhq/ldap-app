FROM alpine:3.8
ADD ./ldap-app /ldap-app
EXPOSE 3000
ENTRYPOINT ["/ldap-app"]
