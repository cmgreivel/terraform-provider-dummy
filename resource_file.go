package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileCreate,
		Read:   resourceFileRead,
		Delete: resourceFileDelete,

		Schema: map[string]*schema.Schema{},
	}
}

func resourceFileCreate(d *schema.ResourceData, m interface{}) error {
	return resourceFileRead(d, m)
}

func resourceFileRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
