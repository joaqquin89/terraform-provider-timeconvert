package timeconvert

import (
    "time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMillisecond() *schema.Resource {
	return &schema.Resource{
		Create: resourceMillisecondCreate,
		Read:   resourceMillisecondRead,
		Update: resourceMillisecondUpdate,
		Delete: resourceMillisecondDelete,
		Schema: map[string]*schema.Schema{
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ms": {
				Type:     schema.TypeInt,
				Computed: true,
				//Sensitive: true,
			},
		},
	}
}

func resourceMillisecondCreate(d *schema.ResourceData, m interface{}) error {
	hours := d.Get("time").(string)
	//ensure that args and scrript is not null
	d.SetId(hours)
	h, _ := time.ParseDuration(hours)
	d.Set("ms", h.Milliseconds())
	return resourceMillisecondRead(d, m)
}

func resourceMillisecondRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceMillisecondUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceMillisecondRead(d, m)
}

func resourceMillisecondDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
