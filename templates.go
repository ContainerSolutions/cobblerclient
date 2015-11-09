package cobblerclient

import (
	"bytes"
	"fmt"
	"io"
)

type loginCredentials struct {
	user string
	pass string
}

func tplLogin(credentials loginCredentials) io.Reader {
	tpl := `<methodCall>
  <methodName>login</methodName>
  <params>
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
  </params>
</methodCall>
`
	txt := fmt.Sprintf(tpl, credentials.user, credentials.pass)
	return bytes.NewReader([]byte(txt))
}

func tplNewSystem(token string) io.Reader {
	tpl := `<methodCall>
  <methodName>new_system</methodName>
  <params>
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
  </params>
</methodCall>
`
	txt := fmt.Sprintf(tpl, token)
	return bytes.NewReader([]byte(txt))
}

func tplSaveSystem(id, token string) io.Reader {
	tpl := `<methodCall>
  <methodName>save_system</methodName>
  <params>
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
  </params>
</methodCall>
`
	txt := fmt.Sprintf(tpl, id, token)
	return bytes.NewReader([]byte(txt))
}

func tplSync(token string) io.Reader {
	tpl := `<methodCall>
  <methodName>sync</methodName>
  <params>
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
  </params>
</methodCall>
`
	txt := fmt.Sprintf(tpl, token)
	return bytes.NewReader([]byte(txt))
}

func tplSetSystemName(id, name, token string) io.Reader {
	tpl := `<param>
      <value>
        <string>%s</string>
      </value>
    </param>`
	txt := fmt.Sprintf(tpl, name)
	return tplModifySystem(id, "name", txt, token)
}

func tplSetSystemProfile(id, profile, token string) io.Reader {
	tpl := `<param>
      <value>
        <string>%s</string>
      </value>
    </param>`
	txt := fmt.Sprintf(tpl, profile)
	return tplModifySystem(id, "profile", txt, token)
}

func tplSetSystemHostname(id, hostname, token string) io.Reader {
	tpl := `<param>
      <value>
        <string>%s</string>
      </value>
    </param>`
	txt := fmt.Sprintf(tpl, hostname)
	return tplModifySystem(id, "hostname", txt, token)
}

func tplSetSystemNameservers(id, nameservers, token string) io.Reader {
	tpl := `<param>
      <value>
        <string>%s</string>
      </value>
    </param>`
	txt := fmt.Sprintf(tpl, nameservers)
	return tplModifySystem(id, "name_servers", txt, token)
}

func tplSetSystemNetwork(id string, config NetworkConfig, token string) io.Reader {
	tpl := `<param>
      <value>
        <struct>
          <member>
            <name>macaddress-eth0</name>
            <value>
              <string>%s</string>
            </value>
          </member>
          <member>
            <name>ipaddress-eth0</name>
            <value>
              <string>%s</string>
            </value>
          </member>
          <member>
            <name>dnsname-eth0</name>
            <value>
              <string>%s</string>
            </value>
          </member>
          <member>
            <name>subnetmask-eth0</name>
            <value>
              <string>%s</string>
            </value>
          </member>
          <member>
            <name>if-gateway-eth0</name>
            <value>
              <string>%s</string>
            </value>
          </member>
        </struct>
      </value>
    </param>`
	txt := fmt.Sprintf(tpl, config.Mac, config.Ip, config.DNSName, config.Netmask, config.Gateway)
	return tplModifySystem(id, "modify_interface", txt, token)
}

// id: systemID
// target: is the target operation that is going to be performed (modify the profile, network, name, what?)
// changes: the XML snippet specific to the change that is going to be yielded in the whole XML request
// token: the authentication/login token to talk to Cobbler.
func tplModifySystem(id, target, changes, token string) io.Reader {
	tpl := `<methodCall>
  <methodName>modify_system</methodName>
  <params>
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
    %s
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
  </params>
</methodCall>
`
	txt := fmt.Sprintf(tpl, id, target, changes, token)
	return bytes.NewReader([]byte(txt))
}

func tplDeleteSystem(name, token string) io.Reader {
	tpl := `<methodCall>
  <methodName>remove_system</methodName>
  <params>
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
    <param>
      <value>
        <string>%s</string>
      </value>
    </param>
  </params>
</methodCall>
`
	txt := fmt.Sprintf(tpl, name, token)
	return bytes.NewReader([]byte(txt))
}
