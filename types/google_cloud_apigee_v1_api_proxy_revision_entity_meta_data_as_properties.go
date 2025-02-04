// Code generated by Sideko (sideko.dev)
package types

import (
	json "encoding/json"
)

type GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties struct {
	AdditionalProperties map[string]string `json:"-"`
}

func (x *GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties) UnmarshalJSON(data []byte) error {
	// this type might have additional properties of type string
	// to handle this scenario we unmarshal twice, once for the named object
	// properties, and once for the additional properties

	// there are no named properties for this type
	additionalPropsBytes := data

	type _GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties // type alias avoids infinite recursion
	var unmarshaled _GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties

	unmarshalErr := json.Unmarshal(data, &unmarshaled)
	if unmarshalErr != nil {
		return unmarshalErr
	}
	err := json.Unmarshal(additionalPropsBytes, &unmarshaled.AdditionalProperties)
	if err != nil {
		return err
	}

	*x = GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties(unmarshaled)

	return nil
}

func (x GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties) MarshalJSON() ([]byte, error) {
	// this type might have additional properties of type string
	// to handle this scenario we must flatten these properties while marshalling
	type _GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties // type alias avoids infinite recursion
	var toMarshal = _GoogleCloudApigeeV1ApiProxyRevisionEntityMetaDataAsProperties(x)

	marshalled, err := json.Marshal(toMarshal)
	if err != nil {
		return nil, err
	}

	var objMap map[string]interface{}
	err = json.Unmarshal(marshalled, &objMap)
	if err != nil {
		return nil, err
	}

	for key, value := range x.AdditionalProperties {
		objMap[key] = value
	}

	return json.Marshal(objMap)
}
