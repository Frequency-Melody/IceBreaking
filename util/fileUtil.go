package util

func IsUploadPicture(contentType string) bool {
	pictureTypes := []string{"image/png", "image/jpeg", "image/gif"}
	return isStringInList(contentType, pictureTypes)
}

func isStringInList(str string, list []string) bool {
	for _, s := range list {
		if str == s {
			return true
		}
	}
	return false
}
