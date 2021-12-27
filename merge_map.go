package merge_map

import "reflect"

func New2Src(src map[string]interface{}, new map[string]interface{}) {
	for k := range new {
		for k1 := range src {
			if k == k1 {
				if reflect.TypeOf(new[k]).Kind() == reflect.Map && reflect.TypeOf(src[k]).Kind() == reflect.Map {
					New2Src(src[k].(map[string]interface{}), new[k].(map[string]interface{}))
				} else {
					src[k1] = new[k]
				}
			}
		}
		src[k] = new[k]
	}
}
func Src2New(src, new map[string]interface{}) {
	for k := range src {
		for k1 := range new {
			if k == k1 {
				if reflect.TypeOf(new[k]).Kind() == reflect.Map && reflect.TypeOf(src[k]).Kind() == reflect.Map {
					Src2New(src[k].(map[string]interface{}), new[k].(map[string]interface{}))
				} else {
					src[k] = new[k]
				}
			}
		}
	}
}
