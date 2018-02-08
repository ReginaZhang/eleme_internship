package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"git.elenet.me/yuelong.huang/pansible/utils"

	"log"
	"net/http"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/franela/goreq"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func printJSON(data interface{}) {
	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func exitOnError(err error, msg string) {
	if err != nil {
		logrus.WithError(err).Fatalln(msg)
	}
}

var debug bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "pansible",
	Short: "A brief description of your application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		m, err := getInventory(viper.GetString("RUN_ID"))
		exitOnError(err, "Pansible failed to get inventory")

		if debug {
			log.Println("get inventory suceeded")
		}

		res, err := getGroup(m.Hosts)
		exitOnError(err, "Failed to parse hosts")

		setVars(res, m.Vars)

		output := utils.Map{
			"_meta": utils.Map{
				"hostvars": utils.Map{},
			},
		}

		for k, v := range res {
			output[k] = v
		}

		j, err := json.MarshalIndent(output, "", "  ")
		exitOnError(err, "Failed to marshal json")

		if debug {
			log.Printf("inv output => %s", string(j))
		}
		log.Println(string(j))
		fmt.Println(string(j))
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	debug = true
	if debug {
		f, err := os.OpenFile("root.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logrus.WithError(err).Fatal("Failed to open log file")
		}
		defer f.Close()

		log.SetOutput(f)
		log.Println("This is a test log entry")
		log.Println("runID => ", viper.GetString("RUN_ID"))
		log.Println("token => ", viper.GetString("token"))

		log.Println("Root execute called")
	}

	if err := RootCmd.Execute(); err != nil {
		if debug {
			log.Println(err)
		}
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./pansible.yml)")
	RootCmd.Flags().Bool("list", false, "")
	RootCmd.Flags().String("host", "", "")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in current directory with name "config" (without extension).
		viper.AddConfigPath(filepath.Dir(os.Args[0]))
		viper.SetConfigName("pansible")
	}

	viper.SetEnvPrefix("PANSIBLE")
	viper.SetDefault("token", "")
	viper.SetDefault("RUN_ID", "")
	viper.SetDefault("MASTER", "")
	viper.SetDefault("SSHKEY_SERVER", "")
	viper.BindPFlag("list", RootCmd.Flag("list"))
	viper.BindPFlag("host", RootCmd.Flag("host"))
	viper.AutomaticEnv() // read in environment variables that match
}

type runInventory struct {
	Hosts utils.Map `json:"hosts"`
	Vars  utils.Map `json:"vars"`
}

func getInventory(runID string) (*runInventory, error) {
	request := goreq.Request{
		//Timeout: time.Second,
		Method: http.MethodGet,
		Uri:    fmt.Sprintf("http://%s/runner/inventory/%s", viper.GetString("MASTER"), runID),
	}

	request.AddHeader(echo.HeaderAuthorization, "Bearer "+viper.GetString("token"))
	response, err := request.Do()
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var resp runInventory
	if err := response.Body.FromJsonTo(&resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

type Group struct {
	Hosts    interface{} `json:"hosts"`
	Vars     utils.Map   `json:"vars,omitempty"`
	Children []string    `json:"children,omitempty"`
}

func getGroup(hi interface{}) (map[string]*Group, error) {
	hosts, err := utils.NewMap(hi)
	if err != nil {
		return nil, errors.Wrap(err, "invalid group type")
	}

	res := make(map[string]*Group)

	for name, gi := range hosts {
		g, err := utils.NewMap(gi)
		if err != nil {
			return nil, errors.Wrap(err, "invalid group type")
		}

		group := Group{
			Hosts: g.Get("hosts"),
		}
		res[name] = &group

		childrenGroups, err := getGroup(g.Get("children"))
		if err != nil {
			return nil, err
		}

		for name, g := range childrenGroups {
			group.Children = append(group.Children, name)
			res[name] = g
		}

	}

	return res, nil
}

func setVars(groups map[string]*Group, vars utils.Map) {
	for name, v := range vars {
		if g, ok := groups[name]; ok {
			g.Vars, _ = utils.NewMap(v)
		}
	}
}
