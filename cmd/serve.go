/*
Copyright © 2021 darmiel <hi@d2a.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/darmiel/yaxc/internal/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:  "serve",
	Long: `Run the YAxC server`,
	Run: func(cmd *cobra.Command, args []string) {
		// create server & start
		s := server.NewServer(&server.YAxCConfig{})
		s.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
