package timeconvert

import (
    "time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceNanosecond() *schema.Resource {
	return &schema.Resource{
		Create: resourceNanosecondCreate,
		Read:   resourceNanosecondRead,
		Update: resourceNanosecondUpdate,
		Delete: resourceNanosecondDelete,
		Schema: map[string]*schema.Schema{
			"time": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ns": {
				Type:     schema.TypeFloat,
				Computed: true,
				//Sensitive: true,
			},
		},
	}
}

func resourceNanosecondCreate(d *schema.ResourceData, m interface{}) error {
	hours := d.Get("time").(string)
	//ensure that args and scrript is not null
	d.SetId(hours)
	h, _ := time.ParseDuration(hours)
	d.Set("ns", h.Nanoseconds())
	return resourceNanosecondRead(d, m)
}

func resourceNanosecondRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceNanosecondUpdate(d *schema.ResourceData, m interface{}) error {

	return resourceNanosecondRead(d, m)
}

func resourceNanosecondDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
