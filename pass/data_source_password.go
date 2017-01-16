package pass

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/hashicorp/terraform/helper/schema"
)

func passwordDataSource() *schema.Resource {
	return &schema.Resource{
		Read: passwordDataSourceRead,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Full path from which a password will be read.",
			},

			"data_row": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String read from Pass.",
			},

			"data": &schema.Schema{
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Map of strings read from Pass.",
			},
		},
	}
}

func passwordDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	path := d.Get("path").(string)

	log.Printf("[DEBUG] Reading %s from Pass", path)
	password, err := exec.Command("pass", path).Output()
	if err != nil {
		return fmt.Errorf("error reading from Pass: %s", err)
	}
	d.Set("data_row", string(password))


	return nil
}
