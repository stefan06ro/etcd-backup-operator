package giantnetes

import "github.com/giantswarm/microerror"

var invalidConfigError = microerror.Error{
	Kind: "invalid config",
}

var failedBackupError = microerror.Error{
	Kind: "backup failed",
}

var unableToGetTenantClustersError = microerror.Error{
	Kind: "unable to get any tenant cluster",
}
