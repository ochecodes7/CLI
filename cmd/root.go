/*
Copyright © 2023 Oche <ochecodes@gmail.com>

*/
package cmd

import (
	"os"
	"fmt"
    "net/http"
    "time"

    "github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	//"golang.org/x/net/dns/dnsutil"
	"github.com/miekg/dns"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "LogQI",
	Short: "LogQI Cli",
	Long: `LogIQ is an open-source web performance monitoring tool that provides comprehensive insights into website performance on both desktop and mobile devices`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	 Run: func(cmd *cobra.Command, args []string) { },
}

rootCmd.Flags().StringVarP(&url, "url", "u", "", "The URL of the web page to monitor")
rootCmd.Flags().IntVarP(&duration, "duration", "d", 10, "The duration in seconds to monitor the web page")

rootCmd.Run = func(cmd *cobra.Command, args []string) {
    start := time.Now()
    client := &http.Client{}
    request, err := http.NewRequest("GET", url, nil)
    if err != nil {
        fmt.Println(err)
        return
    }

    response, err := client.Do(request)
    if err != nil {
        fmt.Println(err)
        return
    }

    duration := time.Since(start).Seconds()
    fmt.Println("The web page took", duration, "seconds to load")

    // Get the response time from the external API endpoint
    responseTime := gjson.Get(response.Body, "responseTime")
    fmt.Println("The response time from the external API endpoint is", responseTime)

    // Analyze the website
    records, err := dns.LookupHost(url)
    if err != nil {
        fmt.Println(err)
        return
    }

    for _, record := range records {
        fmt.Println("Host:", record.Host)
        fmt.Println("IP Address:", record.IP)
    }
}

err := rootCmd.Execute()
if err != nil {
    fmt.Println(err)
}


