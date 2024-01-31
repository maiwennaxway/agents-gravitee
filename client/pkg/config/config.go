package config

import (
	"errors"
	"strings"
	"time"

	"github.com/Axway/agent-sdk/pkg/cmd/properties"
)

type props interface {
	AddStringProperty(name string, defaultVal string, description string)
	AddIntProperty(name string, defaultVal int, description string)
	AddBoolProperty(name string, defaultVal bool, description string)
	AddDurationProperty(name string, defaultVal time.Duration, description string, opts ...properties.DurationOpt)
	StringPropertyValue(name string) string
	IntPropertyValue(name string) int
	BoolPropertyValue(name string) bool
	DurationPropertyValue(name string) time.Duration
}

// GraviteeConfig - represents the config for gateway
type GraviteeConfig struct {
	Http_port                        int
	Http_host                        string
	Http_idleTimeout                 int
	Http_tcpKeepAlive                bool
	Http_compressionSupported        bool
	Http_maxHeaderSize               int
	Http_maxChunkSize                int
	Http_maxInitialLineLength        int
	Http_instances                   int
	Http_requestTimeoutGraceDelay    time.Duration
	Http_secured                     bool
	Http_alpn                        bool
	Http_ssl_client_auth             string
	URL                              string
	APIVersion                       string `config:"apiVersion"`
	EnvID                            string `config:"envID"`
	Man_type                         string
	Ratelimit_type                   string
	Reporter_elasticsearch           bool
	Reporter_elasticsearch_endpoints string
	Reporter_file                    bool
	Services_core                    bool
	Services_core_port               int
	Services_core_host               string
	Auth                             *AuthConfig `config:"auth"`
	Services_sync_delay              time.Duration
	Services_sync_unit               string
	Services_sync_repository         bool
	Services_sync_distributed        bool
	Services_sync_bulkitems          int
	Services_monitoring_delay        time.Duration
	Services_monitoring_unit         string
	Services_monitoring_distributed  bool
	Services_metrics                 bool
	Services_metrics_prometheus      bool
	Services_tracing                 bool
	Services_tracing_type            string
	Ds_mongodb_dbname                string
	Ds_mongodb_host                  string
	Ds_mongodb_port                  int
	Ds_elastic_host                  string
	Ds_elastic_port                  int
	Api_encryption_secret            string
	Classloader_legacy               bool
	CloneAttributes                  bool               `config:"cloneAttributes"`
	Intervals                        *GraviteeIntervals `config:"interval"`
	Workers                          *GraviteeWorkers   `config:"workers"`
	mode                             DiscoveryMode
}

// GraviteeIntervals - intervals for the gravitee agent to use
type GraviteeIntervals struct {
	Proxy   time.Duration `config:"proxy"`
	Spec    time.Duration `config:"spec"`
	Product time.Duration `config:"product"`
	Stats   time.Duration `config:"stats"`
}

// GraviteeWorkers - number of workers for the gravitee agent to use
type GraviteeWorkers struct {
	Proxy   int `config:"proxy"`
	Spec    int `config:"spec"`
	Product int `config:"product"`
}

type DiscoveryMode int

const (
	DiscoveryModeProxy = iota + 1
	DiscoveryModeProduct
)

const (
	DiscoveryModeProxyString   = "proxy"
	DiscoveryModeProductString = "product"
)

func (m DiscoveryMode) String() string {
	return map[DiscoveryMode]string{
		DiscoveryModeProxy:   DiscoveryModeProductString,
		DiscoveryModeProduct: DiscoveryModeProductString,
	}[m]
}

func stringToDiscoveryMode(s string) DiscoveryMode {
	if mode, ok := map[string]DiscoveryMode{
		DiscoveryModeProxyString:   DiscoveryModeProxy,
		DiscoveryModeProductString: DiscoveryModeProduct,
	}[strings.ToLower(s)]; ok {
		return mode
	}
	return 0
}

const (
	pathHttpPort           = "gravitee.Http.port"
	pathHttpHost           = "gravitee.Http.host"
	pathHttpiTo            = "gravitee.Http.idleTimeOut"
	pathHttptcpKA          = "gravitee.Http.tcpKeepAlive"
	pathHttpcs             = "gravitee.Http.compressionsupported"
	pathHttpmaxHeader      = "gravitee.Http.maxHeaderSize"
	pathHttpmaxChunk       = "gravitee.Http.maxChunkSize"
	pathHttpILL            = "gravitee.Http.InitialLineLenght"
	pathHttpInstances      = "gravitee.Http.instances"
	pathHttprTGD           = "gravitee.Http.requestTimeOutGraceDelay"
	pathHttpsecure         = "gravitee.Http.secure"
	pathHttpalpn           = "gravitee.Http.alpn"
	pathHttpsslclientauth  = "gravitee.Http.ssl.clientAuth"
	pathURL                = "gravitee.url"
	pathAPIVersion         = "gravitee.apiVersion"
	pathMantype            = "gravitee.Management.type"
	pathRLtype             = "gravitee.Ratelimit.type"
	pathreporterES         = "gravitee.reporter.elasticsearch"
	pathreporterESEndP     = "gravitee.reporter.elasticsearch.endpoints"
	pathreporterfile       = "gravitee.reporter.file"
	pathServicescore       = "gravite.Services.core"
	pathServicescorehost   = "gravite.Services.core.host"
	pathServicescoreport   = "gravite.Services.core.port"
	pathServicessyncdelay  = "gravitee.Services.sync.delay"
	pathServicessyncunit   = "gravitee.Services.sync.unit"
	pathServicessyncrepo   = "gravitee.Services.sync.repository"
	pathServicessyncdistri = "gravitee.Services.sync.distributed"
	pathServicessyncbulk   = "gravitee.Services.sync.bulk_items"
	pathServicesmondelay   = "gravitee.Services.monitoring.delay"
	pathServicesmonunit    = "gravitee.Services.monitoring.unit"
	pathServicesmondistrib = "gravitee.Services.monitoring.distributed"
	pathServicesmetrics    = "gravitee.Services.metrics"
	pathServicesmetricspro = "gravitee.Services.metrics.prometheus"
	pathServicestracing    = "gravitee.Services.tracing"
	pathServicestracingtyp = "gravitee.Services.tracing.type"
	pathdsmongodbname      = "gravitee.ds.mongodb.dbname"
	pathdsmongodbhost      = "gravitee.ds.mongodb.host"
	pathdsmongodbport      = "gravitee.ds.mongodb.port"
	pathdselastichost      = "gravitee.ds.elastic.host"
	pathdselasticport      = "gravitee.ds.elastic.port"
	pathApiencryption      = "gravitee.Api.encryption.secret"
	pathClassloaderlegacy  = "gravitee.Classloader.legacy"
	pathAuthURL            = "gravitee.auth.url"
	pathAuthServerUsername = "gravitee.auth.serverUsername"
	pathAuthServerPassword = "gravitee.auth.serverPassword"
	pathAuthUsername       = "gravitee.auth.username"
	pathAuthPassword       = "gravitee.auth.password"
	pathAuthtlsP           = "gravitee.auth.tlsProtocols"
	pathAuthtlsC           = "gravitee.auth.tlsCiphers"
	pathAuthkeystoretype   = "gravitee.auth.keystore.type"
	pathAuthkeystorepath   = "gravitee.auth.keystore.path"
	pathAuthkeystorepassd  = "gravitee.auth.keystore.password"
	pathSpecInterval       = "gravitee.interval.spec"
	pathProxyInterval      = "gravitee.interval.proxy"
	pathProductInterval    = "gravitee.interval.product"
	pathStatsInterval      = "gravitee.interval.stats"
	pathenv                = "gravitee.envID"
	pathSpecWorkers        = "gravitee.workers.spec"
	pathProxyWorkers       = "gravitee.workers.proxy"
	pathProductWorkers     = "gravitee.workers.product"
	pathMode               = "gravitee.DiscoveryMode"
	pathCloneAttributes    = "gravitee.cloneAttributes"
)

// AddProperties - adds config needed for gravitee client
func AddProperties(rootProps properties.Properties) {
	rootProps.AddIntProperty(pathHttpPort, 8082, "Port of the Gateway Http Server")
	rootProps.AddStringProperty(pathHttpHost, "0.0.0.0", "Address of the Gateway Http Server")
	rootProps.AddIntProperty(pathHttpiTo, 0, "IdleTimeOut")
	rootProps.AddBoolProperty(pathHttptcpKA, true, "Set to False to Kill Tcp")
	rootProps.AddBoolProperty(pathHttpcs, false, "Set to True to support compression")
	rootProps.AddIntProperty(pathHttpmaxHeader, 8192, "Maximum header size of the Gateway Http Server")
	rootProps.AddIntProperty(pathHttpmaxChunk, 8192, "Maximum chunk size of the Gateway Http Server")
	rootProps.AddIntProperty(pathHttpILL, 4096, "Maximun Initial Line Length of the Gateway Http Server")
	rootProps.AddIntProperty(pathHttpInstances, 1, "Number of instances of the Gateway Http Server")
	rootProps.AddDurationProperty(pathProxyInterval, 30*time.Second, "Delay of the request Timeout Grace")
	rootProps.AddBoolProperty(pathHttpsecure, true, "Set to False to turn off the security")
	rootProps.AddBoolProperty(pathHttpalpn, true, "Set to False to turn off Alpn")
	rootProps.AddStringProperty(pathHttpsslclientauth, "none", "Supports none, request, required") //à check
	rootProps.AddStringProperty(pathMantype, "mongodb", "Repository Type of the Management")
	rootProps.AddStringProperty(pathURL, "https://api.enterprise.gravitee.com", "GRAVITEE Base URL")
	rootProps.AddStringProperty(pathAPIVersion, "v1", "GRAVITEE API Version")
	rootProps.AddStringProperty(pathRLtype, "mongodb", "Repository Type of the Rate Limit")
	rootProps.AddBoolProperty(pathreporterES, true, "Set to false to turn off the elastic search of the reporter")
	rootProps.AddStringProperty(pathreporterESEndP, "Https://${ds.elastic.host}:${ds.elastic.port}", "EndPoint of the elastic search of the reporter")
	rootProps.AddBoolProperty(pathreporterfile, false, "Set to true to turn on the file of the reporter")
	rootProps.AddBoolProperty(pathServicescore, true, "Set to false to disable the service core Http")
	rootProps.AddIntProperty(pathServicescoreport, 18082, "Port of the Http server of the Services core")
	rootProps.AddStringProperty(pathServicescorehost, "localhost", "Host of the Http server of the Services core")
	rootProps.AddDurationProperty(pathServicessyncdelay, 5*time.Second, "Delay of the synchronization of the Services")
	rootProps.AddStringProperty(pathServicessyncunit, "MILLISECONDS", "Unit of the time for the delay")
	rootProps.AddBoolProperty(pathServicessyncrepo, true, "Set to false to disable the sync repository of the Services")
	rootProps.AddBoolProperty(pathServicessyncdistri, false, "Set to true to distribute data synchronization process")
	rootProps.AddIntProperty(pathServicessyncbulk, 100, "Number of items to retrieve durong synchronization")
	rootProps.AddDurationProperty(pathServicesmondelay, 5*time.Second, "Delay of the monitoring of the Services")
	rootProps.AddStringProperty(pathServicesmonunit, "MILLISECONDS", "Unit of the time for the delay")
	rootProps.AddBoolProperty(pathServicesmondistrib, false, "Set to true to distribute data monitoring gathering process")
	rootProps.AddBoolProperty(pathServicesmetrics, false, "Set to true to enable the metrics")
	rootProps.AddBoolProperty(pathServicesmetricspro, true, "Set to false to disable Prometheus metrics")
	rootProps.AddBoolProperty(pathServicestracing, false, "Set to true to enable the tracing of the Services")
	rootProps.AddStringProperty(pathServicestracingtyp, "jaeger", "Type of the service used to traced our Services")
	rootProps.AddStringProperty(pathdsmongodbname, "gravitee", "Name of the database used for ds Services")
	rootProps.AddStringProperty(pathdsmongodbhost, "localhost", "Host of the mongodb datasource")
	rootProps.AddIntProperty(pathdsmongodbport, 27017, "Port of the mongodb datasource")
	rootProps.AddStringProperty(pathdselastichost, "localhost", "Host of the elastic data source")
	rootProps.AddIntProperty(pathdselasticport, 9200, "Port of the elastic data source")
	rootProps.AddStringProperty(pathApiencryption, "vvLJ4Q8Khvv9tm2tIPdkGEdmgKUruAL6", "Encrypt Api properties using this secret")
	rootProps.AddBoolProperty(pathClassloaderlegacy, false, "Set to true to enable the class loader legacy")
	rootProps.AddStringProperty(pathMode, "proxy", "gravitee Organization")
	rootProps.AddStringProperty(pathAuthURL, "Https://login.gravitee.com", "URL to use when authenticating to gravitee")
	rootProps.AddStringProperty(pathAuthServerUsername, "edgecli", "Username to use to when requesting gravitee token")
	rootProps.AddStringProperty(pathAuthServerPassword, "edgeclisecret", "Password to use to when requesting gravitee token")
	rootProps.AddStringProperty(pathAuthUsername, "", "Username to use to authenticate to gravitee")
	rootProps.AddStringProperty(pathAuthPassword, "", "Password for the user to authenticate to gravitee")
	rootProps.AddStringProperty(pathAuthtlsP, "TLSv1.2, TLSv1.3", "Protocols TLS accepted")
	rootProps.AddStringProperty(pathAuthtlsC, "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384, TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384, TLS_ECDHE_RSA_WITH_A", "Ciphers TLS accpeted")
	rootProps.AddStringProperty(pathAuthkeystoretype, "jks", "Supports jks, pem, pkcs12, self-signed")
	rootProps.AddStringProperty(pathAuthkeystorepath, "${gravitee.home}/security/server-keystore.jks", "Path required for certificate's type which is jks")
	rootProps.AddStringProperty(pathAuthkeystorepassd, "Axway123", "Password for KeyStore")
	rootProps.AddBoolProperty(pathCloneAttributes, false, "Set to true to copy the tags when provisioning a Product in product mode")
	rootProps.AddDurationProperty(pathSpecInterval, 30*time.Minute, "The time interval between checking for updated specs", properties.WithLowerLimit(1*time.Minute))
	rootProps.AddDurationProperty(pathProxyInterval, 30*time.Second, "The time interval between checking for updated proxies", properties.WithUpperLimit(5*time.Minute))
	rootProps.AddDurationProperty(pathProductInterval, 30*time.Second, "The time interval between checking for updated products", properties.WithUpperLimit(5*time.Minute))
	rootProps.AddDurationProperty(pathStatsInterval, 5*time.Minute, "The time interval between checking for updated stats", properties.WithLowerLimit(1*time.Minute), properties.WithUpperLimit(15*time.Minute))
	rootProps.AddStringProperty(pathenv, "", "env ID used to create applications")
	rootProps.AddIntProperty(pathProxyWorkers, 10, "Max number of workers discovering proxies")
	rootProps.AddIntProperty(pathSpecWorkers, 20, "Max number of workers discovering specs")
	rootProps.AddIntProperty(pathProductWorkers, 10, "Max number of workers discovering products")
}

// ParseConfig - parse the config on startup
func ParseConfig(rootProps props) *GraviteeConfig {
	return &GraviteeConfig{
		Http_port:                        rootProps.IntPropertyValue(pathHttpPort),
		Http_host:                        rootProps.StringPropertyValue(pathHttpHost),
		Http_idleTimeout:                 rootProps.IntPropertyValue(pathHttpiTo),
		Http_tcpKeepAlive:                rootProps.BoolPropertyValue(pathHttptcpKA),
		Http_compressionSupported:        rootProps.BoolPropertyValue(pathHttpcs),
		Http_maxHeaderSize:               rootProps.IntPropertyValue(pathHttpmaxHeader),
		Http_maxChunkSize:                rootProps.IntPropertyValue(pathHttpmaxChunk),
		Http_maxInitialLineLength:        rootProps.IntPropertyValue(pathHttpILL),
		Http_instances:                   rootProps.IntPropertyValue(pathHttpInstances),
		Http_requestTimeoutGraceDelay:    rootProps.DurationPropertyValue(pathHttprTGD),
		Http_secured:                     rootProps.BoolPropertyValue(pathHttpsecure),
		Http_alpn:                        rootProps.BoolPropertyValue(pathHttpalpn),
		Http_ssl_client_auth:             rootProps.StringPropertyValue(pathHttpsslclientauth),
		URL:                              strings.TrimSuffix(rootProps.StringPropertyValue(pathURL), "/"),
		APIVersion:                       rootProps.StringPropertyValue(pathAPIVersion),
		EnvID:                            rootProps.StringPropertyValue(pathenv),
		Man_type:                         rootProps.StringPropertyValue(pathMantype),
		Ratelimit_type:                   rootProps.StringPropertyValue(pathRLtype),
		Reporter_elasticsearch:           rootProps.BoolPropertyValue(pathreporterES),
		Reporter_elasticsearch_endpoints: strings.TrimSuffix(rootProps.StringPropertyValue(pathreporterESEndP), "/"),
		Reporter_file:                    rootProps.BoolPropertyValue(pathreporterfile),
		Services_core:                    rootProps.BoolPropertyValue(pathServicescore),
		Services_core_port:               rootProps.IntPropertyValue(pathServicescoreport),
		Services_core_host:               rootProps.StringPropertyValue(pathServicescorehost),
		Services_sync_delay:              rootProps.DurationPropertyValue(pathServicessyncdelay),
		Services_sync_unit:               rootProps.StringPropertyValue(pathServicessyncunit),
		Services_sync_repository:         rootProps.BoolPropertyValue(pathServicessyncrepo),
		Services_sync_distributed:        rootProps.BoolPropertyValue(pathServicessyncdistri),
		Services_sync_bulkitems:          rootProps.IntPropertyValue(pathServicessyncbulk),
		Services_monitoring_delay:        rootProps.DurationPropertyValue(pathServicesmondelay),
		Services_monitoring_unit:         rootProps.StringPropertyValue(pathServicesmonunit),
		Services_monitoring_distributed:  rootProps.BoolPropertyValue(pathServicesmondistrib),
		Services_metrics:                 rootProps.BoolPropertyValue(pathServicesmetrics),
		Services_metrics_prometheus:      rootProps.BoolPropertyValue(pathServicesmetricspro),
		Services_tracing:                 rootProps.BoolPropertyValue(pathServicestracing),
		Services_tracing_type:            rootProps.StringPropertyValue(pathServicestracingtyp),
		Ds_mongodb_dbname:                rootProps.StringPropertyValue(pathdsmongodbname),
		Ds_mongodb_host:                  rootProps.StringPropertyValue(pathdsmongodbhost),
		Ds_mongodb_port:                  rootProps.IntPropertyValue(pathdsmongodbport),
		Ds_elastic_host:                  rootProps.StringPropertyValue(pathdselastichost),
		Ds_elastic_port:                  rootProps.IntPropertyValue(pathdselasticport),
		Api_encryption_secret:            rootProps.StringPropertyValue(pathApiencryption),
		Classloader_legacy:               rootProps.BoolPropertyValue(pathClassloaderlegacy),
		mode:                             stringToDiscoveryMode(rootProps.StringPropertyValue(pathMode)),
		CloneAttributes:                  rootProps.BoolPropertyValue(pathCloneAttributes),
		Intervals: &GraviteeIntervals{
			Stats:   rootProps.DurationPropertyValue(pathStatsInterval),
			Proxy:   rootProps.DurationPropertyValue(pathProxyInterval),
			Spec:    rootProps.DurationPropertyValue(pathSpecInterval),
			Product: rootProps.DurationPropertyValue(pathProductInterval),
		},
		Workers: &GraviteeWorkers{
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
			tlsProtocols:   rootProps.StringPropertyValue(pathAuthtlsP),
			tlsCiphers:     rootProps.StringPropertyValue(pathAuthtlsC),
			keystore_type:  rootProps.StringPropertyValue(pathAuthkeystoretype),
			keystore_path:  rootProps.StringPropertyValue(pathAuthkeystorepath),
			keystore_passd: rootProps.StringPropertyValue(pathAuthkeystorepassd),
		},
	}
}

// ValidateCfg - Validates the gateway config
func (a *GraviteeConfig) ValidateCfg() (err error) {
	if a.mode == 0 {
		return errors.New("configuration gravitee non valide: DiscoveryMode doit être proxy ou product")
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
	if a.URL == "" {
		return errors.New("configuration gravitee non valide: url is not configured")
	}
	if a.APIVersion == "" {
		return errors.New("configuration gravitee non valide: api version is not configured")
	}
	if a.EnvID == "" {
		return errors.New("configuration gravitee non valide: env ID is not configured")
	}
	if a.Http_host == "" {
		return errors.New("configuration gravitee non valide: Http_host ne doit pas être une chaîne vide")
	}
	if a.Http_ssl_client_auth != "" {
		return errors.New("configuration gravitee non valide: Http_ssl doit être une chaîne vide")
	}

	if a.Man_type == "" {
		return errors.New("configuration gravitee non valide: Man_type ne doit pas être une chaîne vide")
	}

	if a.Ratelimit_type == "" {
		return errors.New("configuration gravitee non valide: Ratelimit_type ne doit pas être une chaîne vide")
	}

	if a.Reporter_elasticsearch_endpoints == "" {
		return errors.New("configuration gravitee non valide: Reporter_elasticsearch_endpoints ne doit pas être une chaîne vide")
	}

	if a.Services_core_host == "" {
		return errors.New("configuration gravitee non valide: Services_core_host ne doit pas être une chaîne vide")
	}

	if a.Services_sync_unit == "" {
		return errors.New("configuration gravitee non valide: Services_sync_unit ne doit pas être une chaîne vide")
	}

	if a.Services_monitoring_unit == "" {
		return errors.New("configuration gravitee non valide: Services_monitoring_unit ne doit pas être une chaîne vide")
	}

	if a.Services_tracing_type == "" {
		return errors.New("configuration gravitee non valide: Services_tracing_type ne doit pas être une chaîne vide")
	}

	if a.Ds_mongodb_dbname == "" {
		return errors.New("configuration gravitee non valide: Ds_mongodb_dbname ne doit pas être une chaîne vide")
	}

	if a.Ds_mongodb_host == "" {
		return errors.New("configuration gravitee non valide: Ds_mongodb_host ne doit pas être une chaîne vide")
	}

	if a.Ds_elastic_host == "" {
		return errors.New("configuration gravitee non valide: Ds_elastic_host ne doit pas être une chaîne vide")
	}

	if a.Api_encryption_secret == "" {
		return errors.New("configuration gravitee non valide: Api_encryption_secret ne doit pas être une chaîne vide")
	}

	// Ajoutez la validation des champs de type int ici
	if a.Http_port < 1 {
		return errors.New("configuration gravitee non valide: Http_port doit être supérieur à 0")
	}

	if a.Http_idleTimeout < 1 {
		return errors.New("configuration gravitee non valide: Http_idleTimeout doit être supérieur à 0")
	}

	if a.Http_maxHeaderSize < 1 {
		return errors.New("configuration gravitee non valide: Http_maxHeaderSize doit être supérieur à 0")
	}

	if a.Http_maxChunkSize < 1 {
		return errors.New("configuration gravitee non valide: Http_maxChunkSize doit être supérieur à 0")
	}

	if a.Http_maxInitialLineLength < 1 {
		return errors.New("configuration gravitee non valide: Http_maxInitialLineLength doit être supérieur à 0")
	}

	if a.Http_instances < 1 {
		return errors.New("configuration gravitee non valide: Http_instances doit être supérieur à 0")
	}

	if a.Http_requestTimeoutGraceDelay < 1 {
		return errors.New("configuration gravitee non valide: Http_requestTimeoutGraceDelay doit être supérieur à 0")
	}

	if a.Services_core_port < 1 {
		return errors.New("configuration gravitee non valide: Services_core_port doit être supérieur à 0")
	}

	if a.Services_sync_delay < 1 {
		return errors.New("configuration gravitee non valide: Services_sync_delay doit être supérieur à 0")
	}

	if a.Services_sync_bulkitems < 1 {
		return errors.New("configuration gravitee non valide: Services_sync_bulkitems doit être supérieur à 0")
	}

	if a.Services_monitoring_delay < 1 {
		return errors.New("configuration gravitee non valide: Services_monitoring_delay doit être supérieur à 0")
	}

	if a.Ds_mongodb_port < 1 {
		return errors.New("configuration gravitee non valide: Ds_mongodb_port doit être supérieur à 0")
	}

	if a.Ds_elastic_port < 1 {
		return errors.New("configuration gravitee non valide: Ds_elastic_port doit être supérieur à 0")
	}

	// Ajoutez la validation des champs booléens ici
	if a.Services_tracing {
		return errors.New("configuration gravitee non valide: Services_tracing doit être false")
	}

	if a.Classloader_legacy {
		return errors.New("configuration gravitee non valide: Classloader_legacy doit être false")
	}

	if !a.Http_tcpKeepAlive {
		return errors.New("configuration gravitee non valide: Http_tcpKeepAlive doit être true")
	}

	if a.Http_compressionSupported {
		return errors.New("configuration gravitee non valide: Http_compressionSupported doit être false")
	}

	if !a.Http_secured {
		return errors.New("configuration gravitee non valide: Http_secured doit être true")
	}

	if !a.Http_alpn {
		return errors.New("configuration gravitee non valide: Http_alpn doit être true")
	}

	if !a.Reporter_elasticsearch {
		return errors.New("configuration gravitee non valide: Reporter_elasticsearch doit être true")
	}

	if a.Reporter_file {
		return errors.New("configuration gravitee non valide: Reporter_file doit être false")
	}

	if !a.Services_core {
		return errors.New("configuration gravitee non valide: Services_core doit être true")
	}

	if !a.Services_sync_repository {
		return errors.New("configuration gravitee non valide: Services_sync_repository doit être true")
	}

	if a.Services_sync_distributed {
		return errors.New("configuration gravitee non valide: Services_sync_distributed doit être false")
	}

	if a.Services_monitoring_distributed {
		return errors.New("configuration gravitee non valide: Services_monitoring_distributed doit être false")
	}

	if a.Services_metrics {
		return errors.New("configuration gravitee non valide: Services_metrics doit être false")
	}

	if !a.Services_metrics_prometheus {
		return errors.New("configuration gravitee non valide: Services_metrics_prometheus doit être true")
	}

	return

}

// GetAuth - Returns the Auth Config
func (a *GraviteeConfig) GetAuth() *AuthConfig {
	return a.Auth
}

// GetIntervals - Returns the Intervals
func (a *GraviteeConfig) GetIntervals() *GraviteeIntervals {
	return a.Intervals
}

// GetWorkers - Returns the number of Workers
func (a *GraviteeConfig) GetWorkers() *GraviteeWorkers {
	return a.Workers
}

func (a *GraviteeConfig) IsProxyMode() bool {
	return a.mode == DiscoveryModeProxy
}

func (a *GraviteeConfig) IsProductMode() bool {
	return a.mode == DiscoveryModeProduct
}

func (a *GraviteeConfig) ShouldCloneAttributes() bool {
	return a.CloneAttributes
}
