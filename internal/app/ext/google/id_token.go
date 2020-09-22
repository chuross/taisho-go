package google_ext

import (
	"fmt"

	"cloud.google.com/go/compute/metadata"
)

func GetIdToken(serviceID string) (string, error) {
	token, err := metadata.Get(fmt.Sprintf("/instance/service-accounts/default/identity?audience=%s", serviceID))
	if err != nil {
		return "", err
	}
	return token, nil
}
