package cmd

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.hollow.sh/toolbox/ginjwt"

	"go.hollow.sh/dnscontroller/internal/httpsrv"
	dbx "go.hollow.sh/dnscontroller/internal/x/db"
	flagsx "go.hollow.sh/dnscontroller/internal/x/flags"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the dns-controller frontend server",
	Run: func(cmd *cobra.Command, args []string) {
		serve(cmd.Context())
	},
}

func init() {
	root.Cmd.AddCommand(serveCmd)

	logger = root.Options.GetLogger()

	serveCmd.Flags().String("listen", "0.0.0.0:14000", "address on which to listen")
	flagsx.MustBindPFlag("listen", serveCmd.Flags().Lookup("listen"))

	serveCmd.Flags().String("db-uri", "postgresql://root@localhost:26257/dns-controller?sslmode=disable", "URI for database connection")
	flagsx.MustBindPFlag("db.uri", serveCmd.Flags().Lookup("db-uri"))

	flagsx.RegisterOIDCFlags(serveCmd)
}

func serve(ctx context.Context) {
	var db *sqlx.DB
	if viper.GetBool("tracing.enabled") {
		db = dbx.NewDBWithTracing(logger)
	} else {
		db = dbx.NewDB(logger)
	}

	logger.Infow("starting dns-controller frontend server", "address", viper.GetString("listen"))

	hs := &httpsrv.Server{
		Logger: logger,
		Listen: viper.GetString("listen"),
		Debug:  viper.GetBool("logging.debug"),
		DB:     db,
		AuthConfig: ginjwt.AuthConfig{
			Enabled:       viper.GetBool("oidc.enabled"),
			Audience:      viper.GetString("oidc.audience"),
			Issuer:        viper.GetString("oidc.issuer"),
			JWKSURI:       viper.GetString("oidc.jwksuri"),
			LogFields:     viper.GetStringSlice("oidc.log"), // TODO: We don't seem to be grabbing this from config?
			RolesClaim:    viper.GetString("oidc.claims.roles"),
			UsernameClaim: viper.GetString("oidc.claims.username"),
		},
		TrustedProxies: viper.GetStringSlice("gin.trustedproxies"),
	}

	if err := hs.Run(); err != nil {
		logger.Fatalw("failed starting metadata server", "error", err)
	}
}
