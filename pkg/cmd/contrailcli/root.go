package contrailcli

import (
	"strings"

	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/Juniper/contrail/pkg/apisrv/client"
	"github.com/Juniper/contrail/pkg/apisrv/keystone"
	"github.com/Juniper/contrail/pkg/common"
	"github.com/Juniper/contrail/pkg/services"
)

var configFile string

func init() {
	cobra.OnInitialize(initConfig)
	ContrailCLI.PersistentFlags().StringVarP(&configFile, "config", "c", "",
		"Configuration File")
	viper.SetEnvPrefix("contrail")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}

// ContrailCLI defines root Contrail CLI command.
var ContrailCLI = &cobra.Command{
	Use:   "contrailcli",
	Short: "Contrail CLI command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func initConfig() {
	if configFile == "" {
		configFile = viper.GetString("config")
	}
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't read config: ", err)
	}

}

func getClient() (*client.HTTP, error) {
	authURL := viper.GetString("keystone.auth_url")
	client := client.NewHTTP(
		viper.GetString("client.endpoint"),
		authURL,
		viper.GetString("client.id"),
		viper.GetString("client.password"),
		viper.GetString("client.domain_id"),
		viper.GetBool("insecure"),
		&keystone.Scope{
			Project: &keystone.Project{
				Name: viper.GetString("client.project_id"),
				Domain: &keystone.Domain{
					ID: viper.GetString("domain_id"),
				},
			},
		},
	)
	var err error
	if authURL != "" {
		err = client.Login(context.Background()) //nolint
	}
	return client, err
}

// readResources decodes single or array of input data from YAML.
func readResources(file string) (*services.EventList, error) {
	request := &services.EventList{}
	err := common.LoadFile(file, request)
	return request, err
}

func path(schemaID, uuid string) string {
	return "/" + dashedCase(schemaID) + "/" + uuid
}

func pluralPath(schemaID string) string {
	return "/" + dashedCase(schemaID) + "s"
}
