package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

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
	dirName, err := ioutil.TempDir("", "terraform-provider-dummy")

	if err != nil {
		return err
	}

	fullPath := filepath.Join(dirName, "step3.txt")
	data := []byte("Test string\n")

	err = ioutil.WriteFile(fullPath, data, 0644)
	if err != nil {
		return err
	}

	// This provides the ID that TF uses to track this item.
	// If an Id is not set, then TF does not think the object was created.
	d.SetId(fullPath)
	log.Printf("[DEBUG] Created file %s\n", fullPath)
	return resourceFileRead(d, m)
}

func resourceFileRead(d *schema.ResourceData, m interface{}) error {
	data, err := ioutil.ReadFile(d.Id())
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Read '%s' from file %s\n", string(data), d.Id())
	return nil
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
