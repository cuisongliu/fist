// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/fanux/fist/auth"
	"github.com/spf13/cobra"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "For create a kubernetes user jwt token.",
	Run: func(cmd *cobra.Command, args []string) {
		auth.Serve()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)

	// Here you will define your flags and configuration settings.
	authCmd.Flags().Uint16VarP(&auth.AuthHTTPSPort, "https-port", "P", 8443, "start  listening https port")
	authCmd.Flags().Uint16VarP(&auth.AuthHTTPPort, "http-port", "p", 8080, "start  listening http port")
	authCmd.Flags().StringVarP(&auth.AuthCert, "cert", "C", "/etc/fist/cert.pem", "the cert.pem for fist")
	authCmd.Flags().StringVarP(&auth.AuthKey, "key", "K", "/etc/fist/key.pem", "the key.pem for fist")
	authCmd.Flags().StringVarP(&auth.PrivateKey, "pri-pem", "", "/etc/fist/private.pem", "the private.pem for generate key pair")
	authCmd.Flags().StringVarP(&auth.PublicKey, "pub-pem", "", "/etc/fist/public.pem", "the public.pem for generate key pair")
}
