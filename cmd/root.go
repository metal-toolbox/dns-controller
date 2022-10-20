package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go.hollow.sh/toolbox/rootcmd"
	"go.hollow.sh/toolbox/version"
)

const app = "dnscontroller"

var (
	logger *zap.SugaredLogger
)

// root represents the base command when called without any subcommands
var root = rootcmd.NewRootCmd(app, app+" is a tool for centrally managing DNS record from disjoint sources")

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(root.Execute())
}

func init() {
	root.InitFlags()
	cobra.OnInitialize(appInit)
}

func appInit() {
	root.Options.InitConfig()
	logger = setupLogging(root.Options)
}

func setupLogging(o *rootcmd.Options) *zap.SugaredLogger {
	cfg := zap.NewProductionConfig()
	if o.PrettyPrint {
		cfg = zap.NewDevelopmentConfig()
	}

	if o.Debug {
		cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	return l.Sugar().With("app", o.App, "version", version.Version())
}
