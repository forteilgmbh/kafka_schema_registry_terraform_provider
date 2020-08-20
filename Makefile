NEXT_PROVIDER_VERSION=1.1.1

openbsd_amd64:
	CGO_ENABLED=0 GOOS=openbsd GOARCH=amd64 go build -o terraform-provider-schemaregistry_v$(NEXT_PROVIDER_VERSION)
	tar cvzf terraform-provider-schemaregistry_$(NEXT_PROVIDER_VERSION)_openbsd_amd64.tar.gz terraform-provider-schemaregistry_v$(NEXT_PROVIDER_VERSION)

freebsd_amd64:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 go build -o terraform-provider-schemaregistry_v$(NEXT_PROVIDER_VERSION)
	tar cvzf terraform-provider-schemaregistry_$(NEXT_PROVIDER_VERSION)_freebsd_amd64.tar.gz terraform-provider-schemaregistry_v$(NEXT_PROVIDER_VERSION)

linux_amd64:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o terraform-provider-schemaregistry_v$(NEXT_PROVIDER_VERSION)
	tar cvzf terraform-provider-schemaregistry_$(NEXT_PROVIDER_VERSION)_linux_amd64.tar.gz terraform-provider-schemaregistry_v$(NEXT_PROVIDER_VERSION)

darwin_amd64:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o terraform-provider-schemaregistry_v$(NEXT_PROVIDER_VERSION)
	tar cvzf terraform-provider-schemaregistry_$(NEXT_PROVIDER_VERSION)_darwin_amd64.tar.gz terraform-provider-schemaregistry_v$(NEXT_PROVIDER_VERSION)

test:
	go build -o terraform-provider-schemaregistry_v$(NEXT_PROVIDER_VERSION)
	cp terraform-provider-schemaregistry_v$(NEXT_PROVIDER_VERSION) ~/.terraform.d/plugins/
	cd examples; terraform init; terraform apply -auto-approve

clean:
	cd examples; rm -rf .terraform; rm -f *.tfstate*
