package urlshort

import (
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type link struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, req *http.Request) {
		if path, ok := pathsToUrls[req.URL.Path]; ok {
			http.Redirect(w, req, path, http.StatusFound)
		}
		fallback.ServeHTTP(w, req)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:m
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//	yaml := `
// - path: /urlshort
//   url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
//   url: https://github.com/gophercises/urlshort/tree/solution
// `
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	parsedYaml, err := parseYaml(yml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func buildMap(links []link) map[string]string {
	pathMap := make(map[string]string)
	for _, l := range links {
		pathMap[l.Path] = l.URL
	}
	return pathMap
}

func parseYaml(yml []byte) ([]link, error) {
	var links []link
	err := yaml.Unmarshal(yml, &links)
	if err != nil {
		return nil, err
	}
	return links, nil
}
