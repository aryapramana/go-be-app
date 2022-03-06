package commonfile

func GetAPINewsKey() string {
	// put api key here
	return "{{api_key}}"
}

func GetNewsBaseURL() string {
	return "https://newsapi.org/v2/top-headlines"
}

func GetHttpBinBaseURL() string {
	return "https://eu.httpbin.org"
}
