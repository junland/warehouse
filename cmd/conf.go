package cmd

import "github.com/junland/warehouse/server"

func GetEnvConf() server.Config {

	confLogLvl := GetEnvString("WAREHOUSE_LOG_LEVEL", "DEBUG")
	confAccess := GetEnvBool("WAREHOUSE_ENABLE_ACCESS", true)
	confPort := GetEnvString("WAREHOUSE_LISTEN_PORT", "8080")
	confPID := GetEnvString("WAREHOUSE_PID", "/var/run/warehouse.pid")
	confTLS := GetEnvBool("WAREHOUSE_TLS", false)
	confCert := GetEnvString("WAREHOUSE_TLS_CERT", "./cert.pem")
	confKey := GetEnvString("WAREHOUSE_TLS_KEY", "./cert.key")
	confAssetsDir := GetEnvString("WAREHOUSE_ASSETS_DIR", "./")

	config := server.Config{
		LogLvl:    confLogLvl,
		Access:    confAccess,
		Port:      confPort,
		PID:       confPID,
		TLS:       confTLS,
		Cert:      confCert,
		Key:       confKey,
		AssetsDir: confAssetsDir,
	}

	return config

}
