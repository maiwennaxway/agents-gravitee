package gravitee

import (
	"context"
	"fmt"

	//"fmt"

	"github.com/Axway/agent-sdk/pkg/util/log"
	//"github.com/maiwennaxway/agents-gravitee/client/pkg/gravitee/models"
)

func createApiCacheKey(id string) string {
	return fmt.Sprintf("api-%s", id)
}

type ctxKeys string

const (
	loggerKey ctxKeys = "logger"
)

func (c ctxKeys) String() string {
	return string(c)
}

func addLoggerToContext(ctx context.Context, logger log.FieldLogger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func getLoggerFromContext(ctx context.Context) log.FieldLogger {
	return ctx.Value(loggerKey).(log.FieldLogger)
}

func getStringFromContext(ctx context.Context, key ctxKeys) string {
	value := ctx.Value(key)
	if value == nil {
		return "" // ou une autre valeur par défaut selon vos besoins
	}
	if str, ok := value.(string); ok {
		return str
	}
	return ""
}

/*func loadSpecFile(log log.FieldLogger, filePath string) ([]byte, error) {
	log = log.WithField("specFilePath", filePath)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		log.Debug("spec file not found")
		return nil, nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.WithError(err).Error("could not read spec file")
		return nil, err
	}
	return data, nil
}*/
