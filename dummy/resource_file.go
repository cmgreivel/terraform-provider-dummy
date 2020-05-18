package dummy

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileCreate,
		Read:   resourceFileRead,
		Update: resourceFileUpdate,
		Delete: resourceFileDelete,
		Schema: map[string]*schema.Schema{
			"file_name": &schema.Schema {
				Type: schema.TypeString,
				Description: "Name of the file to create",
				Required: true,
				ForceNew: true,
			},
			"contents": &schema.Schema {
				Type: schema.TypeString,
				Description: "File contents",
				Required: true,
			},
		},
	}
}

const filePerm os.FileMode = 0644

func formatCreationLine() string {
	now := time.Now()
	return "CREATION: " + now.Format(time.UnixDate) + "\n"
}

func formatContentsLine(c string) (string, error) {
	if len(c) == 0 {
		return "", errors.New("Cannot have empty string for contents")
	}
	return "CONTENTS: " + c, nil
}

func resourceFileCreate(d *schema.ResourceData, m interface{}) error {

	log.Println("[DEBUG] resourceFileCreate")

	config := m.(Config)

	// Not checking for existence of these keys, because Required is set in the schema.
	fileName := d.Get("file_name").(string)
	contents := d.Get("contents").(string)

	fullPath := filepath.Join(config.Directory, fileName)

	creationLine := formatCreationLine()
	contentLine, err := formatContentsLine(contents)
	if err != nil {
		return err
	}
	data := []byte(creationLine + contentLine)

	err = ioutil.WriteFile(fullPath, data, filePerm)
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
	log.Println("[DEBUG] resourceFileRead")
	data, err := ioutil.ReadFile(d.Id())
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] Read '%s' from file %s\n", string(data), d.Id())
	return nil
}

func resourceFileUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] resourceFileUpdate")
	if d.HasChange("contents") {
		log.Println("[DEBUG] contents changed")
		fileName := d.Id()
		f, err := os.OpenFile(fileName, os.O_RDWR, filePerm)
		if err != nil {
			return err
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		newContents := ""
		for scanner.Scan() {
			line := scanner.Text()
			// Put the newline back on. It was stripped by the scanner.
			line = line + "\n"
			if strings.Contains(line, "CONTENTS:") {
				// Replace the CONTENTS string
				line, err = formatContentsLine(d.Get("contents").(string))
				if err != nil {
					return err
				}
			}
			newContents = newContents + line
		}

		f.Seek(0, io.SeekStart)

		log.Printf("[DEBUG] writing\n\t%s", newContents)
		n, err := f.WriteString(newContents)
		if err != nil {
			return err
		}
		log.Printf("[DEBUG] wrote %d bytes\n", n)
		f.Sync()
	}
	return resourceFileRead(d, m)
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] resourceFileDelete")
	return os.Remove(d.Id())
}
