package cmd

import (
	"context"
	"fmt"
	"github.com/disism/saikan/internal/database"
	"github.com/disism/saikan/internal/nodeinfo"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"text/tabwriter"
)

// oidcCmd represents the oidc command
var oidcCmd = &cobra.Command{
	Use:     "oidc",
	Short:   "OpenID Connect (OIDC)",
	Long:    `saikan oidc-provider configuration command line tools`,
	Aliases: []string{"o"},
}

var addOIDCCmd = &cobra.Command{
	Use:   "add",
	Short: "Add oidc provider supported.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		conf, _ := cmd.Flags().GetString("configuration_endpoint")

		ctx := context.Background()

		client, err := database.New(ctx)
		if err != nil {
			fmt.Println(client)
			return
		}

		if err := nodeinfo.NewClient(&ctx, client.Client).AddOIDC(name, conf); err != nil {
			fmt.Printf("failed to add oidc provider: %v\n", err)
			return
		}
		return
	},
}

var listOIDCCmd = &cobra.Command{
	Use:     "list",
	Short:   "List oidc provider supported.",
	Long:    ``,
	Aliases: []string{"ls"},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		client, err := database.New(ctx)
		if err != nil {
			fmt.Println(client)
			return
		}

		oidc, err := nodeinfo.NewClient(&ctx, client.Client).ListOIDC()
		if err != nil {
			fmt.Println(err)
			return
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.StripEscape)
		fmt.Fprintln(w, "ID\tName\tConfigurationEndpoint")
		for _, o := range oidc {
			fmt.Fprintf(w, "%d\t%s\t%s\n", o.ID, o.Name, o.ConfigurationEndpoint)
		}
		w.Flush()
		return
	},
}

var removeOIDCCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove oidc provider supported.",
	Long:    ``,
	Aliases: []string{"rm"},
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")

		ctx := context.Background()
		client, err := database.New(ctx)
		if err != nil {
			fmt.Println(client)
			return
		}
		p, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			fmt.Println(err)
			return
		}
		if err := nodeinfo.NewClient(&ctx, client.Client).RemoveOIDC(p); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(oidcCmd)
	oidcCmd.AddCommand(addOIDCCmd)
	oidcCmd.AddCommand(listOIDCCmd)
	oidcCmd.AddCommand(removeOIDCCmd)

	addOIDCCmd.Flags().StringP("name", "n", "", "oidc provider name")
	addOIDCCmd.Flags().StringP("configuration_endpoint", "c", "", "oidc provider configuration endpoint")

	removeOIDCCmd.Flags().StringP("id", "i", "", "oidc provider id")

	if err := addOIDCCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
		return
	}

	if err := addOIDCCmd.MarkFlagRequired("configuration_endpoint"); err != nil {
		fmt.Println(err)
		return
	}

	if err := removeOIDCCmd.MarkFlagRequired("id"); err != nil {
		fmt.Println(err)
		return
	}

}
