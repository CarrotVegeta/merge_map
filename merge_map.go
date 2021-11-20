package merge_map

import "reflect"

func New2Src(src map[string]interface{}, new map[string]interface{}) {
	for k := range new {
		for k1 := range src {
			if k == k1 {
				if reflect.TypeOf(new[k]).Kind() == reflect.Map && reflect.TypeOf(src[k]).Kind() == reflect.Map {
					mergeMap(src[k].(map[string]interface{}), new[k].(map[string]interface{}))
				} else {
					src[k1] = new[k]
				}
			}
		}
		src[k] = new[k]
	}
}
