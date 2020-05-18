package dummy

import (
	"errors"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type Config struct {
	Directory string
}

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"directory": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Directory for file creation",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"dummy_file": resourceFile(),
		},
		// Note: This is deprecated should use ConfigureContextFunc
		ConfigureFunc: configure,
	}
}

func configure(d *schema.ResourceData) (interface{}, error) {
	dirName := d.Get("directory").(string)

	stat, err := os.Stat(dirName)
	if err != nil {
		return nil, err
	}
	if ! stat.IsDir() {
		return nil, errors.New(fmt.Sprintf("%s is not a directory\n", dirName))
	}

	return Config{ dirName }, nil
}
