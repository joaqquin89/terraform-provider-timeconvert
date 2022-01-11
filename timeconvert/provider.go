package timeconvert

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"timeconvert_milliseconds": resourceMillisecond(),
			"timeconvert_seconds": resourceSecond(),
			"timeconvert_nanoseconds": resourceNanosecond(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
