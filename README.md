# terraform-provider-dummy

## Introduction
Simple Terraform provider for a tutorial. The provider is a directory, and the provider's 
resource is a file.

Generally follows [Terraform documentation](https://www.terraform.io/docs/extend/writing-custom-providers.html).

## Additional References
- [Terraform plugin SDK docs](https://pkg.go.dev/github.com/hashicorp/terraform-plugin-sdk)
- [Terraform plugin SDK source](https://github.com/hashicorp/terraform-plugin-sdk)
- [Example open source providers](https://github.com/terraform-providers)
- [Posts about writing custom providers](https://github.com/shuaibiyy/awesome-terraform#writing-custom-providers). Note thatthese may be slightly out of date, but the concepts are still applicable.
- [Provider for NetBox](https://github.com/cmgreivel/terraform-provider-netbox)

## Notes
This has all been developed on MacOS. There are not anticipated issues with other OSes.

# Steps

- [x] Create initial buildable version of provider
- [ ] Add resource to the provider
- [ ] Add functioning create and read methods
- [ ] Provider can be configured to specify directory
- [ ] File name and contents can be specified for a resource
- [ ] Unit and acceptance tests added

# Step 1

We have a provider that can be built using
`$ go build -o terraform-provider-dummy`

The binary can be invoked but does nothing.
`$ ./terraform-provider-dummy`

