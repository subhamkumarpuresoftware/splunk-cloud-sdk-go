// Package forwarders -- generated by scloudgen
// !! DO NOT EDIT !!
//
package forwarders

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/flags"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
)

// AddCertificate Adds a certificate to a vacant slot on a tenant.
func AddCertificate(cmd *cobra.Command, args []string) error {

	var err error

	// Parse all flags

	var inputDatafile string
	err = flags.ParseFlag(cmd.Flags(), "input-datafile", &inputDatafile)
	if err != nil {
		return fmt.Errorf(`error parsing "input-datafile": ` + err.Error())
	}

	resp, err := AddCertificateOverride(inputDatafile)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// DeleteCertificate Removes a certificate on a particular slot on a tenant.
func DeleteCertificate(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var slot string
	err = flags.ParseFlag(cmd.Flags(), "slot", &slot)
	if err != nil {
		return fmt.Errorf(`error parsing "slot": ` + err.Error())
	}

	err = client.ForwardersService.DeleteCertificate(slot)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCertificates Removes all certificates on a tenant.
func DeleteCertificates(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	err = client.ForwardersService.DeleteCertificates()
	if err != nil {
		return err
	}

	return nil
}

// ListCertificates Returns a list of all certificates for a tenant.
func ListCertificates(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	resp, err := client.ForwardersService.ListCertificates()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
