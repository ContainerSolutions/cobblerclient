/*
Copyright 2015 Container Solutions

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

package cobblerclient

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
)

// CreateSystemOpts contains fields to set for a new system
type CreateSystemOpts struct {
	BootFiles                string   `mapstructure:"boot_files"`
	Comment                  string   `mapstructure:"comment"`
	EnableGPXE               bool     `mapstructure:"enable_gpxe"`
	FetchableFiles           string   `mapstructure:"fetchable_files"`
	Gateway                  string   `mapstructure:"gateway"`
	Hostname                 string   `mapstructure:"hostname"`
	Image                    string   `mapstructure:"image"`
	IPv6DefaultDevice        string   `mapstructure:"ipv6_default_device"`
	KernelOptions            string   `mapstructure:"kernel_options"`
	KernelOptionsPost        string   `mapstructure:"kernel_options_post"`
	Kickstart                string   `mapstructure:"kickstart"`
	KSMeta                   string   `mapstructure:"ks_meta"`
	LDAPEnabled              bool     `mapstructure:"ldap_enabled"`
	LDAPType                 string   `mapstructure:"ldap_type"`
	MGMTClasses              []string `mapstructure:"mgmt_classes"`
	MGMTParameters           string   `mapstructure:"mgmt_parameters"`
	MonitEnabled             bool     `mapstructure:"monit_enabled"`
	Name                     string   `mapstructure:"name"`
	NameServersSearch        []string `mapstructure:"name_servers_search"`
	NameServers              []string `mapstructure:"name_servers"`
	NetbootEnabled           bool     `mapstructure:"netboot_enabled"`
	Owners                   []string `mapstructure:"owners"`
	PowerAddress             string   `mapstructure:"power_address"`
	PowerID                  string   `mapstructure:"power_id"`
	PowerPass                string   `mapstructure:"power_pass"`
	PowerType                string   `mapstructure:"power_type"`
	PowerUser                string   `mapstructure:"power_user"`
	Profile                  string   `mapstructure:"profile"`
	Proxy                    string   `mapstructure:"proxy"`
	RedHatManagementKey      string   `mapstructure:"redhat_management_key"`
	RedhatManagementServer   string   `mapstructure:"redhat_management_server"`
	Status                   string   `mapstructure:"status"`
	TemplateFiles            string   `mapstructure:"template_files"`
	TemplateRemoteKickstarts int      `mapstructure:"template_remote_kickstarts"`
	VirtAutoBoot             string   `mapstructure:"virt_auto_boot"`
	VirtFileSize             string   `mapstructure:"virt_file_size"`
	VirtCPUs                 string   `mapstructure:"virt_cpus"`
	VirtType                 string   `mapstructure:"virt_type"`
	VirtPath                 string   `mapstructure:"virt_path"`
	VirtPXEBoot              int      `mapstructure:"virt_pxe_boot"`
	VirtRam                  string   `mapstructure:"virt_ram"`
	VirtDiskDriver           string   `mapstructure:"virt_disk_driver"`

	// There can't be a proper Interface struct because of how
	// Cobbler mangles the interface attribute name to "netmask-eth0"
	Interfaces map[string]interface{} `mapstructure:"interfaces"`
}

// System is a created system.
type System struct {
	// These are internal fields and cannot be modified.
	Ctime                 float64 `mapstructure:"ctime"` // TODO: convert to time
	Depth                 int     `mapstructure:"depth"`
	ID                    string  `mapstructure:"uid"`
	IPv6Autoconfiguration bool    `mapstructure:"ipv6_autoconfiguration"`
	Mtime                 float64 `mapstructure:"mtime"` // TODO: convert to time
	ReposEnabled          bool    `mapstructure:"repos_enabled"`

	BootFiles                string                 `mapstructure:"boot_files"`
	Comment                  string                 `mapstructure:"comment"`
	EnableGPXE               bool                   `mapstructure:"enable_gpxe"`
	FetchableFiles           string                 `mapstructure:"fetchable_files"`
	Gateway                  string                 `mapstructure:"gateway"`
	Hostname                 string                 `mapstructure:"hostname"`
	Image                    string                 `mapstructure:"image"`
	Interfaces               map[string]interface{} `mapstructure:"interfaces"`
	IPv6DefaultDevice        string                 `mapstructure:"ipv6_default_device"`
	KernelOptions            string                 `mapstructure:"kernel_options"`
	KernelOptionsPost        string                 `mapstructure:"kernel_options_post"`
	Kickstart                string                 `mapstructure:"kickstart"`
	KSMeta                   string                 `mapstructure:"ks_meta"`
	LDAPEnabled              bool                   `mapstructure:"ldap_enabled"`
	LDAPType                 string                 `mapstructure:"ldap_type"`
	MGMTClasses              []string               `mapstructure:"mgmt_classes"`
	MGMTParameters           string                 `mapstructure:"mgmt_parameters"`
	MonitEnabled             bool                   `mapstructure:"monit_enabled"`
	Name                     string                 `mapstructure:"name"`
	NameServersSearch        []string               `mapstructure:"name_servers_search"`
	NameServers              []string               `mapstructure:"name_servers"`
	NetbootEnabled           bool                   `mapstructure:"netboot_enabled"`
	Owners                   []string               `mapstructure:"owners"`
	PowerAddress             string                 `mapstructure:"power_address"`
	PowerID                  string                 `mapstructure:"power_id"`
	PowerPass                string                 `mapstructure:"power_pass"`
	PowerType                string                 `mapstructure:"power_type"`
	PowerUser                string                 `mapstructure:"power_user"`
	Profile                  string                 `mapstructure:"profile"`
	Proxy                    string                 `mapstructure:"proxy"`
	RedHatManagementKey      string                 `mapstructure:"redhat_management_key"`
	RedhatManagementServer   string                 `mapstructure:"redhat_management_server"`
	Status                   string                 `mapstructure:"status"`
	TemplateFiles            string                 `mapstructure:"template_files"`
	TemplateRemoteKickstarts int                    `mapstructure:"template_remote_kickstarts"`
	VirtAutoBoot             string                 `mapstructure:"virt_auto_boot"`
	VirtFileSize             string                 `mapstructure:"virt_file_size"`
	VirtCPUs                 string                 `mapstructure:"virt_cpus"`
	VirtType                 string                 `mapstructure:"virt_type"`
	VirtPath                 string                 `mapstructure:"virt_path"`
	VirtPXEBoot              int                    `mapstructure:"virt_pxe_boot"`
	VirtRam                  string                 `mapstructure:"virt_ram"`
	VirtDiskDriver           string                 `mapstructure:"virt_disk_driver"`
}

// GetSystems returns all systems in Cobbler.
func (c *Client) GetSystems() ([]System, error) {
	var systems []System

	result, err := c.Call("get_systems", "", c.token)
	if err != nil {
		return nil, err
	}

	for _, s := range result.([]interface{}) {
		var system System
		if err := decodeSystem(s, &system); err != nil {
			return nil, err
		}
		systems = append(systems, system)
	}

	return systems, nil
}

// GetSystem returns a single system obtained by its name.
func (c *Client) GetSystem(name string) (*System, error) {
	var system System

	result, err := c.Call("get_system", name, c.token)
	if err != nil {
		return &system, err
	}

	err = decodeSystem(result, &system)

	return &system, err
}

// CreateSystem creates a system.
// It ensures that either a Profile or Image are set and then sets other default values.
func (c *Client) CreateSystem(system CreateSystemOpts) (*System, error) {
	if system.Profile == "" && system.Image == "" {
		return nil, fmt.Errorf("A system must have a profile or image set.")
	}

	// Set default values. I guess these aren't taken care of by Cobbler?
	if system.BootFiles == "" {
		system.BootFiles = "<<inherit>>"
	}

	if system.FetchableFiles == "" {
		system.FetchableFiles = "<<inherit>>"
	}

	if system.MGMTParameters == "" {
		system.MGMTParameters = "<<inherit>>"
	}

	if system.PowerType == "" {
		system.PowerType = "ipmilan"
	}

	if system.Status == "" {
		system.Status = "production"
	}

	if system.VirtAutoBoot == "" {
		system.VirtAutoBoot = "0"
	}

	if system.VirtCPUs == "" {
		system.VirtCPUs = "<<inherit>>"
	}

	if system.VirtDiskDriver == "" {
		system.VirtDiskDriver = "<<inherit>>"
	}

	if system.VirtFileSize == "" {
		system.VirtFileSize = "<<inherit>>"
	}

	if system.VirtPath == "" {
		system.VirtPath = "<<inherit>>"
	}

	if system.VirtRam == "" {
		system.VirtRam = "<<inherit>>"
	}

	if system.VirtType == "" {
		system.VirtType = "<<inherit>>"
	}

	// To create a system via the Cobbler API, first call new_system to obtain an ID
	result, err := c.Call("new_system", c.token)
	if err != nil {
		return nil, err
	}
	newId := result.(string)

	// Cobbler wants each field to be updated individually.
	// Also, for some reason, arrays are flattened to space-delimited strings.
	// Finally, Interface options are passed by a proper struct.
	s := reflect.ValueOf(&system).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		tag := typeOfT.Field(i).Tag
		cobblerField := tag.Get("mapstructure")

		switch f.Type().String() {
		case "string", "bool", "int64", "int":
			if err := c.UpdateSystemField(newId, cobblerField, f.Interface()); err != nil {
				return nil, err
			}
		case "[]string":
			v := strings.Join(f.Interface().([]string), " ")
			if err := c.UpdateSystemField(newId, cobblerField, v); err != nil {
				return nil, err
			}
		case "map[string]interface {}":
			for nicName, nicData := range system.Interfaces {
				nic := map[string]interface{}{}
				for k, v := range nicData.(map[string]interface{}) {
					attrName := fmt.Sprintf("%s-%s", k, nicName)
					nic[attrName] = v
				}
				if err := c.UpdateSystemField(newId, "modify_interface", nic); err != nil {
					return nil, err
				}
			}
		default:
			fmt.Printf("%s\n", f.Type().String())
		}
	}

	if _, err := c.Call("save_system", newId, c.token); err != nil {
		return nil, err
	}

	return c.GetSystem(system.Name)
}

// UpdateSystemField updates a single field in a given system.
func (c *Client) UpdateSystemField(systemId, field, value interface{}) error {
	if result, err := c.Call("modify_system", systemId, field, value, c.token); err != nil {
		return err
	} else {
		if result.(bool) == false {
			return fmt.Errorf("Error updating %s to %s.", field, value)
		}
	}

	return nil
}

// decodeSystem is a custom mapstructure decoder to handle Cobbler's uniqueness.
func decodeSystem(raw interface{}, system *System) error {
	var metadata mapstructure.Metadata
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:         &metadata,
		Result:           system,
		WeaklyTypedInput: true,
		DecodeHook:       systemDataHacks,
	})

	if err != nil {
		return err
	}

	if err := decoder.Decode(raw); err != nil {
		return err
	}

	return nil
}

// systemDataHacks is a hook for the mapstructure decoder.
// It's used to smooth out issues with converting fields and types from Cobbler.
func systemDataHacks(f, t reflect.Kind, data interface{}) (interface{}, error) {
	dataVal := reflect.ValueOf(data)
	if dataVal.String() == "~" {
		return map[string]interface{}{}, nil
	}
	if f == reflect.Int64 && t == reflect.Bool {
		if dataVal.Int() > 0 {
			return true, nil
		} else {
			return false, nil
		}
	}
	return data, nil
}
