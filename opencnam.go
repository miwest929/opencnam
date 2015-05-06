// Uses http://www.opencnam.com/
// CONTRIBUTING: go install && opencnam query <phone-number>

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "OpenCNAM Querier"
	app.Usage = "Provide a phone number and the Caller ID will be returned"
	app.Commands = []cli.Command{
		{
			Name:      "query",
			ShortName: "q",
			Usage:     "Query for the Caller ID",
			Action: func(c *cli.Context) {
				phoneNumber := c.Args().First()

				println("Caller ID for ", phoneNumber)

				parts := []string{"https://api.opencnam.com/v2/phone", phoneNumber}
				resp, err := http.Get(strings.Join(parts, "/"))
				if err != nil {
					println("Error retrieving Caller ID")
				} else {
					defer resp.Body.Close()
					body, _ := ioutil.ReadAll(resp.Body)

					fmt.Printf("%s\n", string(body))
				}
			},
		},
	}

	app.Run(os.Args)
}
