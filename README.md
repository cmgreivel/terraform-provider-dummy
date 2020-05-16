# terraform-provider-dummy
Simple Terraform provider for a tutorial.

Generally follows [Terraform documentation](https://www.terraform.io/docs/extend/writing-custom-providers.html)

# Steps

- [ ] Create initial buildable version of provider
- [ ] Add resource to the provider
- [ ] Add functioning create and read methods
- [ ] Provider can be configured to specify directory
- [ ] File name and contents can be specified for a resource

# Step 1

We have a provider that can be built using
`$ go build -o terraform-provider-dummy`

The binary can be invoked but does nothing.
`$ ./terraform-provider-dummy`

# Step 2

The provider now includes a resource, but still does nothing.

# Step 3

We should now be able to actually use this provider in Terraform and
see some results.

```
$ go build -o terraform-provider-dummy
$ cp terraform-provider-dummy ~/.terraform.d/plugins/darwin_amd64/
```

The above example is for Mac.
* For Linux use `~/.terraform.d/plugins/linux_amd64/`
* For Windows use `%APPDATA%\terraform.d\plugins\windows_amd64`

Alternatively, we can specify the path to the provider on the command line when doing `terraform init`.

```
$ terraform init -plugin-dir=`pwd`
$ terraform validate
$ terraform plan
$ terraform apply
```

The output of the apply should have the path to the file we created as the ID.
```
dummy_file.my_file: Creating...
dummy_file.my_file: Creation complete after 0s [id=/var/folders/d9/0fpnfyr91k5_7pl5mqd5sjc80000gn/T/terraform-provider-dummy644603600/step3.txt]
```

We can verify the contents are as expected.

# Step 4

We can configure the provider to use a specific directory.

We have also rearranged the source directory structure for better maintainability.

Try `terraform validate`, `terraform plan`, and `terraform apply` with different
values for `directory` in `main.tf`. You can try a non-existent directory name or
the name of a file to see what happens.
```
provider dummy {
  directory = "/tmp"
}
```

# Step 5

The file name and contents can be specified as input variables to the resource.
Changing `contents` in `main.tf` will only modify the existing file. Changing `file_name`
will destroy an existing resource and create a new one.

```
resource dummy_file my_file {
  file_name = "file.tmp"
  contents  = "initial contents\n"
}
```

