package customjson

import (
	jsoniter "github.com/json-iterator/go"
)

const DGRAPH = "dgraph"

func getJsonWithCustomTag(tagName string) jsoniter.API {
	return jsoniter.Config{TagKey: tagName}.Froze()
}

func Marshal(object interface{}) string {
	s, _ := jsoniter.MarshalToString(object)
	return s
}

func MarshalWithCustomTag(tagName string, object interface{}) string {
	json := getJsonWithCustomTag(tagName)
	s, _ := json.MarshalToString(object)
	return s
}

func UnmarshalWithCustomTag(tagName, jsonStr string, object interface{}) error {
	json := getJsonWithCustomTag(tagName)
	return json.UnmarshalFromString(jsonStr, object)
}

func UnmarshalFromString(jsonStr string, object interface{}) error {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	return json.UnmarshalFromString(jsonStr, object)
}

func ToDgraphJson(object interface{}) string {
	return MarshalWithCustomTag(DGRAPH, object)
}

func FromDgraphJson(jsonBytes []byte, object interface{}) error {
	json := getJsonWithCustomTag(DGRAPH)
	return json.Unmarshal(jsonBytes, object)
}
