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
	"github.com/fanux/fist/rbac"
	"github.com/spf13/cobra"
)

// rbacCmd represents the rbac command
var rbacCmd = &cobra.Command{
	Use:   "rbac",
	Short: "rbac for palm.",
	Run: func(cmd *cobra.Command, args []string) {
		rbac.Serve()
	},
}

func init() {
	rootCmd.AddCommand(rbacCmd)

	// Here you will define your flags and configuration settings.
	rbacCmd.Flags().Uint16VarP(&rbac.RbacPort, "port", "P", 8080, "start  listening port")

	rbacCmd.Flags().BoolVarP(&rbac.RbacLdapEnable, "ldap-enable", "", false, "ldap enable config.")
	rbacCmd.Flags().Uint16VarP(&rbac.RbacLdapPort, "ldap-port", "", 389, "ldap port config.  e.g. 389")
	rbacCmd.Flags().StringVarP(&rbac.RbacLdapHost, "ldap-host", "", "ldap.com", "ldap host config. e.g. mydomain.com")
	rbacCmd.Flags().StringVarP(&rbac.RbacLdapBindDN, "ldap-bind-dn", "", "", "ldap  bind-dn config. e.g. cn=Search,dc=mydomain,dc=com")
	rbacCmd.Flags().StringVarP(&rbac.RbacLdapBindPassword, "ldap-bind-password", "", "", "ldap bind-password config.")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rbacCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rbacCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
