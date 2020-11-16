// Package api provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/6xRTY/TQAz9KyOfQya0t7mV1R4qAVqpR8TBTNxmIBkPY5eqRPnvyCFIFXDcW+J5fh9+",
	"M0SeCmfKKhDmpYGUzwxhhp4k1lQ0cYYAByc4lZHclxF/Dt+4VYqDO7wcoYExRcpCtpRxIkMXjAO5XdtB",
	"A9c6QoBBtUjw/na7tbg+t1wvftsV//749Pzx9Pxm13btoNMISwOadDS6/2j+oCq/nb1tu7YzNBfKWBIE",
	"2K+jBgrqYKnAY0k+claMav+FRf/NeKLcux4VnbKbSAQv5AbM/UgVVv6KBj32EOCFRQ8lPW2kDVT6fiXR",
	"d9zfjdrUKK8qWMqY4rrqv4pJzSBxoAntS+/FMorWlC+wLEvzl60PmxPlzcwmlir1ELReabGBFLZDGuWu",
	"617Xwh92d648ue2Q7sx1MrThhapVAuHT/FB48P6xu7Dv9p11Acvn5VcAAAD//3A81LV6AgAA",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}