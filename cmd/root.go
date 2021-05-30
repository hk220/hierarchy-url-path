/*
Copyright © 2021 Kazuki Hara

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
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hierarchy-url-path",
	Short: "Test hirerachy url path",
	Run: func(cmd *cobra.Command, args []string) {
		r := mux.NewRouter()
		r.StrictSlash(true)
		s := &Server{}

		r.HandleFunc("/ipxe", s.stage1).Methods("GET")
		r.HandleFunc("/server/{serial:[A-Z0-9]+}/ipxe", s.stage2).Methods("GET")

		r.Use(func(h http.Handler) http.Handler {
			return handlers.LoggingHandler(os.Stdout, h)
		})
		r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))
		http.ListenAndServe(":"+viper.GetString("port"), r)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().IntP("port", "p", 5121, "port number")
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
}
