package swagger

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/hashicorp/go-multierror"
	"github.com/solo-io/gloo-api/pkg/api/types/v1"
	"github.com/solo-io/gloo-function-discovery/internal/detector"
	"github.com/solo-io/gloo-function-discovery/internal/updater/swagger"
	"github.com/solo-io/gloo-plugins/rest"
	"github.com/solo-io/gloo/pkg/log"
)

var commonSwaggerURIs = []string{
	"/swagger.json",
	"/swagger/docs/v1",
	"/swagger/docs/v2",
	"/v1/swagger",
	"/v2/swagger",
}

type swaggerDetector struct {
	swaggerUrisToTry []string
}

func NewSwaggerDetector(swaggerUrisToTry []string) detector.Interface {
	return &swaggerDetector{
		swaggerUrisToTry: append(commonSwaggerURIs, swaggerUrisToTry...),
	}
}

func (d *swaggerDetector) DetectFunctionalService(addr string) (*v1.ServiceInfo, map[string]string, error) {
	var errs error
	for _, uri := range d.swaggerUrisToTry {
		url := "http://" + addr + uri
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return nil, nil, errors.Wrap(err, "invalid url for request")
		}
		req.Header.Set("X-Gloo-Discovery", "Swagger-Discovery")
		res, err := http.Get(url)
		if err != nil {
			errs = multierror.Append(errs, errors.Wrapf(err, "could not perform HTTP GET on resolved addr: %v", addr))
			continue
		}
		// might have found a swagger service
		if res.StatusCode == 200 {
			if _, err := swagger.RetrieveSwaggerDocFromUrl(url); err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			// definitely found swagger
			log.Printf("swagger upstream detected: %v", addr)
			svcInfo := &v1.ServiceInfo{
				Type: rest.ServiceTypeREST,
			}
			annotations := map[string]string{swagger.AnnotationKeySwaggerURL: url}
			return svcInfo, annotations, nil
		} else {
			errs = multierror.Append(errs, errors.Errorf("path: %v response code: %v headers: %v", uri, res.Status, res.Header))
		}
	}
	// not a swagger upstream
	return nil, nil, errors.Wrapf(errs, "service at %s does not implement swagger at a known endpoint, "+
		"or was unreachable", addr)
}
