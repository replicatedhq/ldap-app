
docker:
	GOOS=linux go build .
	docker build -t ldap-app:0.1 .
