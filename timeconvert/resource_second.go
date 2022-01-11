package timeconvert

import (
    "time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSecond() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecondCreate,
		Read:   resourceSecondRead,
		Update: resourceSecondUpdate,
		Delete: resourceSecondDelete,
		Schema: map[string]*schema.Schema{
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"secs": {
				Type:     schema.TypeFloat,
				Computed: true,
				//Sensitive: true,
			},
		},
	}
}

func resourceSecondCreate(d *schema.ResourceData, m interface{}) error {
	hours := d.Get("time").(string)
	//ensure that args and scrript is not null
	d.SetId(hours)
	h, _ := time.ParseDuration(hours)
	d.Set("secs", h.Seconds())
	return resourceSecondRead(d, m)
}

func resourceSecondRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceSecondUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceSecondRead(d, m)
}

func resourceSecondDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
