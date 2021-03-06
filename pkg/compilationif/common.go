package compilationif

import (
	"context"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/Juniper/contrail/pkg/db/etcd"
	"github.com/Juniper/contrail/pkg/services"
)

// NewCompilationService makes a compilation service.
func NewCompilationService() *CompilationService {
	service := &CompilationService{
		BaseService: services.BaseService{},
	}
	Init()
	return service
}

// CompilationService implements Intent Compilation layer.
type CompilationService struct {
	services.BaseService
}

// HandleEtcdMessage handles messages received from etcd.
func (service *CompilationService) HandleEtcdMessage(ctx context.Context, oper int32, key, value string) {
	messageFields := log.Fields{"operation": oper, "key": key, "value": value}
	log.WithFields(messageFields).Print("HandleEtcdMessages: Got a message")
	err := service.handleEtcdMessage(ctx, oper, key, value)
	if err != nil {
		log.WithFields(messageFields).WithError(err).Error("Failed to handle etcd message")
	}
}

func (service *CompilationService) handleEtcdMessage(ctx context.Context, oper int32, key, value string) error {
	processor := &services.ServiceEventProcessor{Service: service}
	event, err := etcd.ParseEvent(oper, key, []byte(value))
	if err != nil {
		return errors.Wrap(err, "failed to parse ETCD event")
	}

	_, err = processor.Process(ctx, event)
	return err
}
