package url_shortener

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if dest, ok := pathsToUrls[request.URL.Path]; ok {
			http.Redirect(writer, request, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(writer, request)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlBytes)

	if err != nil {
		return nil, err
	}

	pathsToUrls := buildMap(pathUrls)

	return MapHandler(pathsToUrls, fallback), nil
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl

	err := yaml.Unmarshal(data, &pathUrls)

	if err != nil {
		return nil, err
	}

	return pathUrls, nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathToUrls := make(map[string]string)

	for _, path := range pathUrls {
		pathToUrls[path.Path] = path.URL
	}

	return pathToUrls
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
