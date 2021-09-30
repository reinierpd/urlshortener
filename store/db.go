package store

var shortUrlsDb = map[string]string{}

func AddUrl(key string, url string) {
	shortUrlsDb[key] = url
}

func GetLongUrl(key string) string {
	return shortUrlsDb[key]
}
