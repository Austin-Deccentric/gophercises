package urlshort

import (
	"encoding/json"
	"log"
	"net/http"
	yaml "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path];  ok {
			log.Println("Redirecting")
			http.Redirect(w,r,dest, http.StatusFound)
			return
		}
		log.Println("Url not found. back to home")
		fallback.ServeHTTP(w,r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	parsedYaml,err := parseYaml(yml); if err != nil {
		return nil,err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}


//ParseYaml is a yaml parser
func parseYaml(data []byte) ([]pathURL, error){
	pathurls := make([]pathURL, len(data))
	if err:= yaml.Unmarshal(data, &pathurls); err != nil {
		return nil, err
	}
	return pathurls, nil

}

func buildMap(inputData []pathURL) map[string]string{
	urlMap := make(map[string]string)
	for _, data := range inputData {
		urlMap[data.Path] = data.URL
	}
	return urlMap
}

//JSONHandler takes a slice of byte of  valid JSON item.
//and a fallback handler that is to be called if URL path is not found.
// NB: Valid JSON item:
// {
// 	"path": <URL Path>,
// 	"url": "<The URL to be accessed>"
// }
//Example:
// {
// 	"path": "/bolt",
// 	"url": "https://github.com/boltdb/bolt"
// } 
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error){
	var pathUrls []pathURL
	if err:= json.Unmarshal(jsn, &pathUrls); err != nil{
		return nil, err
	}
	pathsToUrls := buildMap(pathUrls)
	return MapHandler(pathsToUrls, fallback), nil
}




type pathURL struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

