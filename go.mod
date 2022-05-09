module go-learning

go 1.15

require (
	github.com/aws/aws-sdk-go v1.34.9
	github.com/gofrs/flock v0.8.0
	github.com/goinaction/code v0.0.0-20171020164608-49fc99e6affb
	github.com/pkg/errors v0.9.1
	github.com/varstr/uaparser v0.0.0-20170929040706-6aabb7c4e98c
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	gopkg.in/yaml.v2 v2.4.0
	helm.sh/helm/v3 v3.7.0
	k8s.io/apimachinery v0.22.1
	k8s.io/client-go v0.22.1
	github.com/robfig/cron/v3 v3.0.0
)

replace (
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v17.12.0-ce-rc1.0.20200618181300-9dc6525e6118+incompatible
)
