package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

const (
	COVID_ALL     = "https://corona.lmao.ninja/all"
	COVID_COUNTRY = "https://corona.lmao.ninja/countries/"
)

// CovidCases represents total Covid cases across all countries
type CovidCases struct {
	Region   string `json:"country,omitempty"`
	Cases     int    `json:"cases"`
	Active    int    `json:"active"`
	Recovered int    `json:"recovered"`
	Deaths    int    `json:"deaths"`
}

func displayData(format string, data []CovidCases, limit int) error {
	if limit == 0 {
		limit = 1
	}

	switch format {
	case "json":
		data, err := json.MarshalIndent(data, "", " ")
		if err != nil {
			return err
		}
		fmt.Println(string(data))
	case "table":
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Region", "Cases", "Active", "Recovered", "Deaths"})
		for i, row := range data {
			if i >= limit {
				break
			}
			t.AppendRow([]interface{}{row.Region, row.Cases, row.Active, row.Recovered, row.Deaths})
		}
		t.Style().Options.SeparateRows = true
		t.Style().Title.Align = text.AlignCenter
		configs := []table.ColumnConfig{}
		configs = append(configs, table.ColumnConfig{
			Number:5, 
			Colors: text.Colors{text.FgHiRed, text.BgBlack},
		})
		configs = append(configs, table.ColumnConfig{
			Number:4, 
			Colors: text.Colors{text.FgHiGreen, text.BgBlack},
		})
		configs = append(configs, table.ColumnConfig{
			Number:3, 
			Colors: text.Colors{text.FgHiYellow, text.BgBlack},
		})
		t.SetColumnConfigs(configs)
		t.Render()
	}
	return nil
}

func globalCases(cmd *cobra.Command, args []string) error {

	covidAPI := COVID_ALL
	if len(args) > 0 {
		value := args[0]
		if value == "full" {
			covidAPI = COVID_COUNTRY + "?sort=cases"
		} else {
			covidAPI = COVID_COUNTRY + value
		}
	}

	resp, err := http.Get(covidAPI)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var AllCases []CovidCases
	if err := json.Unmarshal(responseBody, &AllCases); err != nil {
		var stats CovidCases
		stats.Region = "Global"
		err := json.Unmarshal(responseBody, &stats)
		if err != nil {
			return err
		}
		AllCases = append(AllCases, stats)
	}

	format, _ := cmd.Flags().GetString("format")
	limit, _ := cmd.Flags().GetInt("limit")
	displayData(format, AllCases, limit)
	return nil
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "covid",
	Short: "Brings covid19 stats from the comfort of your terminal",
	Long: `Covid is easy to use CLI tool for live covid stats.
	
	Run:
	> covid - Get Global count.

	Subcommands
	> covid full - Get Top counts by active cases
	> covid <country> - Get stats for input country

	Example :
	covid India
	+--------+-------+--------+-----------+--------+
	| REGION | CASES | ACTIVE | RECOVERED | DEATHS |
	+--------+-------+--------+-----------+--------+
	| India  |  1251 |   1117 |       102 |     32 |
	+--------+-------+--------+-----------+--------+
	`,
	// Run : func(cmd *cobra.Command, args []string) {	}
	RunE: globalCases,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.Flags().String("format", "table", "Format for displaying output. Options are Table and Json")
	RootCmd.Flags().Int("limit", 50, "Limit results to display for command <covid full>. Only applicable for table format. Minimum value is 1.")
}
