// Package register provides ...
package register

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
)

type Registrar interface {
	Register(serviceIP string, servicePort int, serviceName string, logger log.Logger) (sd.Registrar, error)
}
