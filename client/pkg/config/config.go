package config

import (
	"errors"
	"strings"
	"time"

	"github.com/Axway/agent-sdk/pkg/cmd/properties"
)

// graviteeConfig - represents the config for gateway
type graviteeConfig struct {
	http_port                        int
	http_host                        string
	http_idleTimeout                 int
	http_tcpKeepAlive                bool
	http_compressionSupported        bool
	http_maxHeaderSize               int
	http_maxChunkSize                int
	http_maxInitialLineLength        int
	http_instances                   int
	http_requestTimeoutGraceDelay    time.Duration
	http_secured                     bool
	http_alpn                        bool
	http_ssl                         string
	man_type                         string
	ratelimit_type                   string
	reporter_elasticsearch           bool
	reporter_elasticsearch_endpoints string
	reporter_file                    bool
	services_core                    bool
	services_core_port               int
	services_core_host               string
	Auth                             *AuthConfig `config:"auth"`
	services_sync_delay              time.Duration
	services_sync_unit               string
	services_sync_repository         bool
	services_sync_distributed        bool
	services_sync_bulkitems          int
	services_monitoring_delay        time.Duration
	services_monitoring_unit         string
	services_monitoring_distributed  bool
	services_metrics                 bool
	services_metrics_prometheus      bool
	services_tracing                 bool
	services_tracing_type            string
	ds_mongodb_dbname                string
	ds_mongodb_host                  string
	ds_mongodb_port                  int
	ds_elastic_host                  string
	ds_elastic_port                  int
	api_encryption_secret            string
	classloader_legacy               bool

	Intervals *graviteeIntervals `config:"interval"`
	Workers   *graviteeWorkers   `config:"workers"`
	mode      discoveryMode
}

// graviteeIntervals - intervals for the gravitee agent to use
type graviteeIntervals struct {
	Proxy   time.Duration `config:"proxy"`
	Spec    time.Duration `config:"spec"`
	Product time.Duration `config:"product"`
	Stats   time.Duration `config:"stats"`
}

// graviteeWorkers - number of workers for the gravitee agent to use
type graviteeWorkers struct {
	Proxy   int `config:"proxy"`
	Spec    int `config:"spec"`
	Product int `config:"product"`
}

type discoveryMode int

const (
	discoveryModeProxy = iota + 1
	discoveryModeProduct
)

const (
	discoveryModeProxyString   = "proxy"
	discoveryModeProductString = "product"
)

func (m discoveryMode) String() string {
	return map[discoveryMode]string{
		discoveryModeProxy:   discoveryModeProductString,
		discoveryModeProduct: discoveryModeProductString,
	}[m]
}

func stringToDiscoveryMode(s string) discoveryMode {
	if mode, ok := map[string]discoveryMode{
		discoveryModeProxyString:   discoveryModeProxy,
		discoveryModeProductString: discoveryModeProduct,
	}[strings.ToLower(s)]; ok {
		return mode
	}
	return 0
}

const (
	pathHttpPort           = "gravitee.http.port"
	pathHttpHost           = "gravitee.http.host"
	pathHttpiTo            = "gravitee.http.idleTimeOut"
	pathHttptcpKA          = "gravitee.http.tcpKeepAlive"
	pathHttpcs             = "gravitee.http.compressionsupported"
	pathHttpmaxHeader      = "gravitee.http.maxHeaderSize"
	pathHttpmaxChunk       = "gravitee.http.maxChunkSize"
	pathHttpILL            = "gravitee.http.InitialLineLenght"
	pathHttpInstances      = "gravitee.http.instances"
	pathHttprTGD           = "gravitee.http.requestTimeOutGraceDelay"
	pathHttpsecure         = "gravitee.http.secure"
	pathHttpalpn           = "gravitee.http.alpn"
	pathHttpssl            = "gravitee.http.ssl"
	pathmantype            = "gravitee.management.type"
	pathRLtype             = "gravitee.ratelimit.type"
	pathreporterES         = "gravitee.reporter.elasticsearch"
	pathreporterESEndP     = "gravitee.reporter.elasticsearch.endpoints"
	pathreporterfile       = "gravitee.reporter.file"
	pathservicescore       = "gravite.services.core"
	pathservicescorehost   = "gravite.services.core.host"
	pathservicescoreport   = "gravite.services.core.port"
	pathservicessyncdelay  = "gravitee.services.sync.delay"
	pathservicessyncunit   = "gravitee.services.sync.unit"
	pathservicessyncrepo   = "gravitee.services.sync.repository"
	pathservicessyncdistri = "gravitee.services.sync.distributed"
	pathservicessyncbulk   = "gravitee.services.sync.bulk_items"
	pathservicesmondelay   = "gravitee.services.monitoring.delay"
	pathservicesmonunit    = "gravitee.services.monitoring.unit"
	pathservicesmondistrib = "gravitee.services.monitoring.distributed"
	pathservicesmetrics    = "gravitee.services.metrics"
	pathservicesmetricspro = "gravitee.services.metrics.prometheus"
	pathservicestracing    = "gravitee.services.tracing"
	pathservicestracingtyp = "gravitee.services.tracing.type"
	pathdsmongodbname      = "gravitee.ds.mongodb.dbname"
	pathdsmongodbhost      = "gravitee.ds.mongodb.host"
	pathdsmongodbport      = "gravitee.ds.mongodb.port"
	pathdselastichost      = "gravitee.ds.elastic.host"
	pathdselasticport      = "gravitee.ds.elastic.port"
	pathapiencryption      = "gravitee.api.encryption.secret"
	pathclassloaderlegacy  = "gravitee.classloader.legacy"
	pathAuthURL            = "gravitee.auth.url"
	pathAuthServerUsername = "gravitee.auth.serverUsername"
	pathAuthServerPassword = "gravitee.auth.serverPassword"
	pathAuthUsername       = "gravitee.auth.username"
	pathAuthPassword       = "gravitee.auth.password"
	pathSpecInterval       = "gravitee.interval.spec"
	pathProxyInterval      = "gravitee.interval.proxy"
	pathProductInterval    = "gravitee.interval.product"
	pathStatsInterval      = "gravitee.interval.stats"
	pathDeveloper          = "gravitee.developerID"
	pathSpecWorkers        = "gravitee.workers.spec"
	pathProxyWorkers       = "gravitee.workers.proxy"
	pathProductWorkers     = "gravitee.workers.product"
	pathMode               = "gravitee.discoveryMode"
)

// AddProperties - adds config needed for gravitee client
func AddProperties(rootProps properties.Properties) {
	rootProps.AddIntProperty(pathHttpPort, 8082, "Port of the Gateway HTTP Server")
	rootProps.AddStringProperty(pathHttpHost, "0.0.0.0", "Address of the Gateway Http Server")
	rootProps.AddIntProperty(pathHttpiTo, 0, "IdleTimeOut")
	rootProps.AddBoolProperty(pathHttptcpKA, true, "Set to False to Kill Tcp")
	rootProps.AddBoolProperty(pathHttpcs, false, "Set to True to support compression")
	rootProps.AddIntProperty(pathHttpmaxHeader, 8192, "Maximum header size of the Gateway HTTP Server")
	rootProps.AddIntProperty(pathHttpmaxChunk, 8192, "Maximum chunk size of the Gateway HTTP Server")
	rootProps.AddIntProperty(pathHttpILL, 4096, "Maximun Initial Line Length of the Gateway HTTP Server")
	rootProps.AddIntProperty(pathHttpInstances, 1, "Number of instances of the Gateway HTTP Server")
	rootProps.AddDurationProperty(pathProxyInterval, 30*time.Second, "Delay of the request Timeout Grace")
	rootProps.AddBoolProperty(pathHttpsecure, true, "Set to False to turn off the security")
	rootProps.AddBoolProperty(pathHttpalpn, true, "Set to False to turn off Alpn")
	rootProps.AddStringProperty(pathHttpssl, "", "") //à check
	rootProps.AddStringProperty(pathmantype, "mongodb", "Repository Type of the Management")
	rootProps.AddStringProperty(pathRLtype, "mongodb", "Repository Type of the Rate Limit")
	rootProps.AddBoolProperty(pathreporterES, true, "Set to false to turn off the elastic search of the reporter")
	rootProps.AddStringProperty(pathreporterESEndP, "https://${ds.elastic.host}:${ds.elastic.port}", "EndPoint of the elastic search of the reporter")
	rootProps.AddBoolProperty(pathreporterfile, false, "Set to true to turn on the file of the reporter")
	rootProps.AddBoolProperty(pathservicescore, true, "Set to false to disable the service core http")
	rootProps.AddIntProperty(pathservicescoreport, 18082, "Port of the Http server of the services core")
	rootProps.AddStringProperty(pathservicescorehost, "localhost", "Host of the Http server of the services core")
	rootProps.AddDurationProperty(pathservicessyncdelay, 5*time.Second, "Delay of the synchronization of the services")
	rootProps.AddStringProperty(pathservicessyncunit, "MILLISECONDS", "Unit of the time for the delay")
	rootProps.AddBoolProperty(pathservicessyncrepo, true, "Set to false to disable the sync repository of the services")
	rootProps.AddBoolProperty(pathservicessyncdistri, false, "Set to true to distribute data synchronization process")
	rootProps.AddIntProperty(pathservicessyncbulk, 100, "Number of items to retrieve durong synchronization")
	rootProps.AddDurationProperty(pathservicesmondelay, 5*time.Second, "Delay of the monitoring of the services")
	rootProps.AddStringProperty(pathservicesmonunit, "MILLISECONDS", "Unit of the time for the delay")
	rootProps.AddBoolProperty(pathservicesmondistrib, false, "Set to true to distribute data monitoring gathering process")
	rootProps.AddBoolProperty(pathservicesmetrics, false, "Set to true to enable the metrics")
	rootProps.AddBoolProperty(pathservicesmetricspro, true, "Set to false to disable Prometheus metrics")
	rootProps.AddBoolProperty(pathservicestracing, false, "Set to true to enable the tracing of the services")
	rootProps.AddStringProperty(pathservicestracingtyp, "jaeger", "Type of the service used to traced our services")
	rootProps.AddStringProperty(pathdsmongodbname, "gravitee", "Name of the database used for ds services")
	rootProps.AddStringProperty(pathdsmongodbhost, "localhost", "Host of the mongodb datasource")
	rootProps.AddIntProperty(pathdsmongodbport, 27017, "Port of the mongodb datasource")
	rootProps.AddStringProperty(pathdselastichost, "localhost", "Host of the elastic data source")
	rootProps.AddIntProperty(pathdselasticport, 9200, "Port of the elastic data source")
	rootProps.AddStringProperty(pathapiencryption, "vvLJ4Q8Khvv9tm2tIPdkGEdmgKUruAL6", "Encrypt API properties using this secret")
	rootProps.AddBoolProperty(pathclassloaderlegacy, false, "Set to true to enable the class loader legacy")
	rootProps.AddStringProperty(pathMode, "proxy", "gravitee Organization")
	rootProps.AddStringProperty(pathAuthURL, "https://login.gravitee.com", "URL to use when authenticating to gravitee")
	rootProps.AddStringProperty(pathAuthServerUsername, "edgecli", "Username to use to when requesting gravitee token")
	rootProps.AddStringProperty(pathAuthServerPassword, "edgeclisecret", "Password to use to when requesting gravitee token")
	rootProps.AddStringProperty(pathAuthUsername, "", "Username to use to authenticate to gravitee")
	rootProps.AddStringProperty(pathAuthPassword, "", "Password for the user to authenticate to gravitee")
	rootProps.AddDurationProperty(pathSpecInterval, 30*time.Minute, "The time interval between checking for updated specs", properties.WithLowerLimit(1*time.Minute))
	rootProps.AddDurationProperty(pathProxyInterval, 30*time.Second, "The time interval between checking for updated proxies", properties.WithUpperLimit(5*time.Minute))
	rootProps.AddDurationProperty(pathProductInterval, 30*time.Second, "The time interval between checking for updated products", properties.WithUpperLimit(5*time.Minute))
	rootProps.AddDurationProperty(pathStatsInterval, 5*time.Minute, "The time interval between checking for updated stats", properties.WithLowerLimit(1*time.Minute), properties.WithUpperLimit(15*time.Minute))
	rootProps.AddStringProperty(pathDeveloper, "", "Developer ID used to create applications")
	rootProps.AddIntProperty(pathProxyWorkers, 10, "Max number of workers discovering proxies")
	rootProps.AddIntProperty(pathSpecWorkers, 20, "Max number of workers discovering specs")
	rootProps.AddIntProperty(pathProductWorkers, 10, "Max number of workers discovering products")
}

// ParseConfig - parse the config on startup
func ParseConfig(rootProps properties.Properties) *graviteeConfig {
	return &graviteeConfig{
		http_port:                        rootProps.IntPropertyValue(pathHttpPort),
		http_host:                        rootProps.StringPropertyValue(pathHttpHost),
		http_idleTimeout:                 rootProps.IntPropertyValue(pathHttpiTo),
		http_tcpKeepAlive:                rootProps.BoolPropertyValue(pathHttptcpKA),
		http_compressionSupported:        rootProps.BoolPropertyValue(pathHttpcs),
		http_maxHeaderSize:               rootProps.IntPropertyValue(pathHttpmaxHeader),
		http_maxChunkSize:                rootProps.IntPropertyValue(pathHttpmaxChunk),
		http_maxInitialLineLength:        rootProps.IntPropertyValue(pathHttpILL),
		http_instances:                   rootProps.IntPropertyValue(pathHttpInstances),
		http_requestTimeoutGraceDelay:    rootProps.IntPropertyValue(pathHttprTGD),
		http_secured:                     rootProps.BoolPropertyValue(pathHttpsecure),
		http_alpn:                        rootProps.BoolPropertyValue(pathHttpalpn),
		http_ssl:                         rootProps.StringPropertyValue(pathHttpssl),
		man_type:                         rootProps.StringPropertyValue(pathmantype),
		ratelimit_type:                   rootProps.StringPropertyValue(pathRLtype),
		reporter_elasticsearch:           rootProps.BoolPropertyValue(pathreporterES),
		reporter_elasticsearch_endpoints: strings.TrimSuffix(rootProps.StringPropertyValue(pathreporterESEndP), "/"),
		reporter_file:                    rootProps.BoolPropertyValue(pathreporterfile),
		services_core:                    rootProps.BoolPropertyValue(pathservicescore),
		services_core_port:               rootProps.IntPropertyValue(pathservicescoreport),
		services_core_host:               rootProps.StringPropertyValue(pathservicescorehost),
		services_sync_delay:              rootProps.DurationPropertyValue(pathservicessyncdelay),
		services_sync_unit:               rootProps.StringPropertyValue(pathservicessyncunit),
		services_sync_repository:         rootProps.BoolPropertyValue(pathservicessyncrepo),
		services_sync_distributed:        rootProps.BoolPropertyValue(pathservicessyncdistri),
		services_sync_bulkitems:          rootProps.IntPropertyValue(pathservicessyncbulk),
		services_monitoring_delay:        rootProps.DurationPropertyValue(pathservicesmondelay),
		services_monitoring_unit:         rootProps.StringPropertyValue(pathservicesmonunit),
		services_monitoring_distributed:  rootProps.BoolPropertyValue(pathservicesmondistrib),
		services_metrics:                 rootProps.BoolPropertyValue(pathservicesmetrics),
		services_metrics_prometheus:      rootProps.BoolPropertyValue(pathservicesmetricspro),
		services_tracing:                 rootProps.BoolPropertyValue(pathservicestracing),
		services_tracing_type:            rootProps.StringPropertyValue(pathservicestracingtyp),
		ds_mongodb_dbname:                rootProps.StringPropertyValue(pathdsmongodbname),
		ds_mongodb_host:                  rootProps.StringPropertyValue(pathdsmongodbhost),
		ds_mongodb_port:                  rootProps.IntPropertyValue(pathdsmongodbport),
		ds_elastic_host:                  rootProps.StringPropertyValue(pathdselastichost),
		ds_elastic_port:                  rootProps.IntPropertyValue(pathdselasticport),
		api_encryption_secret:            rootProps.StringPropertyValue(pathapiencryption),
		classloader_legacy:               rootProps.BoolPropertyValue(pathclassloaderlegacy),
		mode:                             stringToDiscoveryMode(rootProps.StringPropertyValue(pathMode)),
		Intervals: &graviteeIntervals{
			Stats:   rootProps.DurationPropertyValue(pathStatsInterval),
			Proxy:   rootProps.DurationPropertyValue(pathProxyInterval),
			Spec:    rootProps.DurationPropertyValue(pathSpecInterval),
			Product: rootProps.DurationPropertyValue(pathProductInterval),
		},
		Workers: &graviteeWorkers{
			Proxy:   rootProps.IntPropertyValue(pathProxyWorkers),
			Spec:    rootProps.IntPropertyValue(pathSpecWorkers),
			Product: rootProps.IntPropertyValue(pathProductWorkers),
		},
		Auth: &AuthConfig{
			Username:       rootProps.StringPropertyValue(pathAuthUsername),
			Password:       rootProps.StringPropertyValue(pathAuthPassword),
			ServerUsername: rootProps.StringPropertyValue(pathAuthServerUsername),
			ServerPassword: rootProps.StringPropertyValue(pathAuthServerPassword),
			URL:            rootProps.StringPropertyValue(pathAuthURL),
		},
	}
}

// ValidateCfg - Validates the gateway config
func (a *graviteeConfig) ValidateCfg() (err error) {
	if a.mode == 0 {
		return errors.New("configuration gravitee non valide: discoveryMode doit être proxy ou product")
	}

	if a.Auth.Username == "" {
		return errors.New("configuration gravitee non valide: le nom d'utilisateur n'est pas configuré")
	}

	if a.Auth.Password == "" {
		return errors.New("configuration gravitee non valide: le mot de passe n'est pas configuré")
	}

	if a.Workers.Proxy < 1 {
		return errors.New("configuration gravitee non valide: les travailleurs proxy doivent être supérieurs à 0")
	}

	if a.Workers.Spec < 1 {
		return errors.New("configuration gravitee non valide: les travailleurs spec doivent être supérieurs à 0")
	}
	// Ajoutez la validation des champs de type string ici
	if a.http_host == "" {
		return errors.New("configuration gravitee non valide: http_host ne doit pas être une chaîne vide")
	}
	if a.http_ssl != "" {
		return errors.New("configuration gravitee non valide: http_ssl doit être une chaîne vide")
	}

	if a.man_type == "" {
		return errors.New("configuration gravitee non valide: man_type ne doit pas être une chaîne vide")
	}

	if a.ratelimit_type == "" {
		return errors.New("configuration gravitee non valide: ratelimit_type ne doit pas être une chaîne vide")
	}

	if a.reporter_elasticsearch_endpoints == "" {
		return errors.New("configuration gravitee non valide: reporter_elasticsearch_endpoints ne doit pas être une chaîne vide")
	}

	if a.services_core_host == "" {
		return errors.New("configuration gravitee non valide: services_core_host ne doit pas être une chaîne vide")
	}

	if a.services_sync_unit == "" {
		return errors.New("configuration gravitee non valide: services_sync_unit ne doit pas être une chaîne vide")
	}

	if a.services_monitoring_unit == "" {
		return errors.New("configuration gravitee non valide: services_monitoring_unit ne doit pas être une chaîne vide")
	}

	if a.services_tracing_type == "" {
		return errors.New("configuration gravitee non valide: services_tracing_type ne doit pas être une chaîne vide")
	}

	if a.ds_mongodb_dbname == "" {
		return errors.New("configuration gravitee non valide: ds_mongodb_dbname ne doit pas être une chaîne vide")
	}

	if a.ds_mongodb_host == "" {
		return errors.New("configuration gravitee non valide: ds_mongodb_host ne doit pas être une chaîne vide")
	}

	if a.ds_elastic_host == "" {
		return errors.New("configuration gravitee non valide: ds_elastic_host ne doit pas être une chaîne vide")
	}

	if a.api_encryption_secret == "" {
		return errors.New("configuration gravitee non valide: api_encryption_secret ne doit pas être une chaîne vide")
	}

	// Ajoutez la validation des champs de type int ici
	if a.http_port < 1 {
		return errors.New("configuration gravitee non valide: http_port doit être supérieur à 0")
	}

	if a.http_idleTimeout < 1 {
		return errors.New("configuration gravitee non valide: http_idleTimeout doit être supérieur à 0")
	}

	if a.http_maxHeaderSize < 1 {
		return errors.New("configuration gravitee non valide: http_maxHeaderSize doit être supérieur à 0")
	}

	if a.http_maxChunkSize < 1 {
		return errors.New("configuration gravitee non valide: http_maxChunkSize doit être supérieur à 0")
	}

	if a.http_maxInitialLineLength < 1 {
		return errors.New("configuration gravitee non valide: http_maxInitialLineLength doit être supérieur à 0")
	}

	if a.http_instances < 1 {
		return errors.New("configuration gravitee non valide: http_instances doit être supérieur à 0")
	}

	if a.http_requestTimeoutGraceDelay < 1 {
		return errors.New("configuration gravitee non valide: http_requestTimeoutGraceDelay doit être supérieur à 0")
	}

	if a.services_core_port < 1 {
		return errors.New("configuration gravitee non valide: services_core_port doit être supérieur à 0")
	}

	if a.services_sync_delay < 1 {
		return errors.New("configuration gravitee non valide: services_sync_delay doit être supérieur à 0")
	}

	if a.services_sync_bulkitems < 1 {
		return errors.New("configuration gravitee non valide: services_sync_bulkitems doit être supérieur à 0")
	}

	if a.services_monitoring_delay < 1 {
		return errors.New("configuration gravitee non valide: services_monitoring_delay doit être supérieur à 0")
	}

	if a.ds_mongodb_port < 1 {
		return errors.New("configuration gravitee non valide: ds_mongodb_port doit être supérieur à 0")
	}

	if a.ds_elastic_port < 1 {
		return errors.New("configuration gravitee non valide: ds_elastic_port doit être supérieur à 0")
	}

	// Ajoutez la validation des champs booléens ici
	if a.services_tracing {
		return errors.New("configuration gravitee non valide: services_tracing doit être false")
	}

	if a.classloader_legacy {
		return errors.New("configuration gravitee non valide: classloader_legacy doit être false")
	}

	if !a.http_tcpKeepAlive {
		return errors.New("configuration gravitee non valide: http_tcpKeepAlive doit être true")
	}

	if a.http_compressionSupported {
		return errors.New("configuration gravitee non valide: http_compressionSupported doit être false")
	}

	if !a.http_secured {
		return errors.New("configuration gravitee non valide: http_secured doit être true")
	}

	if !a.http_alpn {
		return errors.New("configuration gravitee non valide: http_alpn doit être true")
	}

	if !a.reporter_elasticsearch {
		return errors.New("configuration gravitee non valide: reporter_elasticsearch doit être true")
	}

	if a.reporter_file {
		return errors.New("configuration gravitee non valide: reporter_file doit être false")
	}

	if !a.services_core {
		return errors.New("configuration gravitee non valide: services_core doit être true")
	}

	if !a.services_sync_repository {
		return errors.New("configuration gravitee non valide: services_sync_repository doit être true")
	}

	if a.services_sync_distributed {
		return errors.New("configuration gravitee non valide: services_sync_distributed doit être false")
	}

	if a.services_monitoring_distributed {
		return errors.New("configuration gravitee non valide: services_monitoring_distributed doit être false")
	}

	if a.services_metrics {
		return errors.New("configuration gravitee non valide: services_metrics doit être false")
	}

	if !a.services_metrics_prometheus {
		return errors.New("configuration gravitee non valide: services_metrics_prometheus doit être true")
	}

	return

}

// GetAuth - Returns the Auth Config
func (a *graviteeConfig) GetAuth() *AuthConfig {
	return a.Auth
}

// GetIntervals - Returns the Intervals
func (a *graviteeConfig) GetIntervals() *graviteeIntervals {
	return a.Intervals
}

// GetWorkers - Returns the number of Workers
func (a *graviteeConfig) GetWorkers() *graviteeWorkers {
	return a.Workers
}

func (a *graviteeConfig) IsProxyMode() bool {
	return a.mode == discoveryModeProxy
}

func (a *graviteeConfig) IsProductMode() bool {
	return a.mode == discoveryModeProduct
}
