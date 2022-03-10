package clusters

import "fmt"

const defaultApiVersion = "2020-03-20"

func userAgent() string {
	return fmt.Sprintf("pandora/clusters/%s", defaultApiVersion)
}
