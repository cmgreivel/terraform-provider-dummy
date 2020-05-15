# terraform-provider-dummy
Simple Terraform provider for a tutorial

Generally follows [Terraform documentation](https://www.terraform.io/docs/extend/writing-custom-providers.html)

# Steps

- [ ] Create initial buildable version of compiler
- [ ] Add resource to the provider

# Step 1

We have a provider that can be built using
`$ go build -o terraform-provider-dummy`

The binary can be invoked but does nothing.
`$ ./terraform-provider-dummy`

# Step 2

The provider now includes a resource, but still does nothing.
