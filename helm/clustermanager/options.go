package clustermanager

type ClusterOptions struct {
	*ClusterConfig
	ID         string
	AccountID  int64
	Kubeconfig []byte
	CaCrt      []byte
	CaKey      []byte
	Apiserver  string
}

type ClusterConfig struct {
	EnableExternalAccess bool
	EnableClusterMetrics bool
	UsingCacheClient     bool
	UsingCacheClientStat bool
	IndexerFuncs         []IndexerFunc
	CacheObject          []client.Object
	CheckerConfig        *HealthCheckConfig
	TrackingTime         time.Duration
	KubeTimeout          time.Duration
	KubeQPS              int
	KubeBurst            int
}
