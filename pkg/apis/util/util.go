package util

import "strings"

const (
	// `latest` is a special version that always points to the latest version of the image.
	VersionLatest = "latest"
)

// GetImageVersion returns the verion of a image
func GetImageVersion(image string) string {
	colonIdx := strings.LastIndexByte(image, ':')
	if colonIdx >= 0 {
		return image[colonIdx+1:]
	}

	return VersionLatest
}
