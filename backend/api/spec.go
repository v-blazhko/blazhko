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

	"H4sIAAAAAAAC/5SSwW7bMAyGX0XgdvQsL7vplgU9BNiGAj0WPagyHauTJU1iFnSB332gbK9pnMtugsif",
	"/El+ZzBhiMGjpwzqDNn0OOjy3DmLnnbBkzbEHzGFiIksljAO2jp+tJhNspFs8KDm7wroNSIoyJSsP8BY",
	"gdP+sE6PCTtMCVvB8aM+4C3tgDlzaCVfAjc0Xg83BGYaR5ToSjVWkPDX0SZsQT3CnLSMtDR7qoAsORZO",
	"OxJLVRZjprfC4fkFDcHIla3vwtrQVmQ9RIfi2ek//c9QE5pebO/3UIGzBn0uU0zTwDZq06PY1A1UcEwO",
	"FPREMSspT6dTrUu4DukgZ22W3/a7ux8Pd582dVP3NDjezWL/Rs/fmPLk7HPd1A1nh4heRwsKvpSvCqKm",
	"vkAgdbTSXCASMq1nfEDfilaTFhTEvEXRa986TFDqJ82p+xYU3IdM22gX7qaLYKavoX3l0twNfemiY3TW",
	"FKl8ydxqAZhfHxN2oOCDfCNcznjL92yX87y3/H12SWE2CpdoUDpiYSXHwEvmdpum+S971+itLCzVRZfC",
	"8A+xLqSBszk/Y+JzgXo8X8FweVi+EYxP498AAAD//ywtydXrAwAA",
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
