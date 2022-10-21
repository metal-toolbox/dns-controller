// Package flags contains functions shared
package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// MustBindPFlag binds a viper name to a spf13/pflag and panics on error
func MustBindPFlag(name string, flag *pflag.Flag) {
	err := viper.BindPFlag(name, flag)
	if err != nil {
		panic(err)
	}
}

// RegisterOIDCFlags has flags common to OIDC
func RegisterOIDCFlags(cmd *cobra.Command) {
	cmd.Flags().Bool("oidc", true, "use oidc auth")
	MustBindPFlag("oidc.enabled", cmd.Flags().Lookup("oidc"))
	cmd.Flags().String("oidc-aud", "", "expected audience on OIDC JWT")
	MustBindPFlag("oidc.audience", cmd.Flags().Lookup("oidc-aud"))
	cmd.Flags().StringSlice("oidc-issuer", []string{}, "expected issuer of OIDC JWT")
	MustBindPFlag("oidc.issuer", cmd.Flags().Lookup("oidc-issuer"))
	cmd.Flags().StringSlice("oidc-jwksuri", []string{}, "URI for JWKS listing for JWTs")
	MustBindPFlag("oidc.jwksuri", cmd.Flags().Lookup("oidc-jwksuri"))
	cmd.Flags().String("oidc-roles-claim", "claim", "field containing the permissions of an OIDC JWT")
	MustBindPFlag("oidc.claims.roles", cmd.Flags().Lookup("oidc-roles-claim"))
	cmd.Flags().String("oidc-username-claim", "", "additional fields to output in logs from the JWT token, ex (email)")
	MustBindPFlag("oidc.claims.username", cmd.Flags().Lookup("oidc-username-claim"))
}
