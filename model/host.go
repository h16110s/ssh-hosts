package model

import (
	"reflect"
	"strings"
)

type Host struct {
	Num                              int
	Host                             string
	Match                            string
	AddressFamily                    string
	BatchMode                        string
	BindAddress                      string
	ChallengeResponeAuthentication   string
	ChackHostIP                      string
	Cipher                           string
	Ciphers                          string
	ClearAllForwardings              string
	Complression                     string
	CompresionLevel                  string
	ConnectionAttempts               string
	ConnectTimeout                   string
	ControlMaster                    string
	ControlPath                      string
	DynamicForward                   string
	EscapeChar                       string
	ExitOnForwardFailure             string
	ForwardAgent                     string
	ForwardX11                       string
	ForwardX11Trusted                string
	GatewayPorts                     string
	GlobalKnownHostsFile             string
	GSSAPIAuthentication             string
	GSSAPIKeyExchange                string
	GSSAPIClientIdentity             string
	GSSAPIDelegateCredentials        string
	GSSAPIRenewalForcesRekey         string
	GSSAPITrustDns                   string
	HashKnownHosts                   string
	HostbasedAuthentication          string
	HostKeyAlgorithms                string
	HostKeyAlias                     string
	HostName                         string
	IdentitiesOnly                   string
	IdentityFile                     string
	KbdInteractiveAuthentication     string
	KbdInteractiveDevices            string
	LocalCommand                     string
	LocalForward                     string
	LogLevel                         string
	MACs                             string
	NoHostAuthenticationForLocalhost string
	PreferredAuthentications         string
	Protocol                         string
	ProxyCommand                     string
	PubkeyAuthentication             string
	RemoteForward                    string
	RhostsRSAAuthentication          string
	RSAAuthentication                string
	SendEnv                          string
	ServerAliveCountMax              string
	ServerAliveInterval              string
	SmartcardDevice                  string
	StrictHostKeyChecking            string
	TCPKeepAlive                     string
	Tunnel                           string
	TunnelDevice                     string
	UsePrivilegedPort                string
	UserKnownHostsFile               string
	VerifyHostKeyDNS                 string
	VisualHostKey                    string
}

func (h *Host) Setter(key string, value string) {
	switch strings.ToUpper(key) {
	case "MATCH":
		h.Match = value
	case "ADDRESSFAMILY":
		h.AddressFamily = value
	case "BATCHMODE":
		h.BatchMode = value
	case "BINDADDRESS":
		h.BindAddress = value
	case "CHALLENGERESPONEAUTHENTICATION":
		h.ChallengeResponeAuthentication = value
	case "CHACKHOSTIP":
		h.ChackHostIP = value
	case "CIPHER":
		h.Cipher = value
	case "CIPHERS":
		h.Ciphers = value
	case "CLEARALLFORWARDINGS":
		h.ClearAllForwardings = value
	case "COMPLRESSION":
		h.Complression = value
	case "COMPRESIONLEVEL":
		h.CompresionLevel = value
	case "CONNECTIONATTEMPTS":
		h.ConnectionAttempts = value
	case "CONNECTTIMEOUT":
		h.ConnectTimeout = value
	case "CONTROLMASTER":
		h.ControlMaster = value
	case "CONTROLPATH":
		h.ControlPath = value
	case "DYNAMICFORWARD":
		h.DynamicForward = value
	case "ESCAPECHAR":
		h.EscapeChar = value
	case "EXITONFORWARDFAILURE":
		h.ExitOnForwardFailure = value
	case "FORWARDAGENT":
		h.ForwardAgent = value
	case "FORWARDX11":
		h.ForwardX11 = value
	case "FORWARDX11TRUSTED":
		h.ForwardX11Trusted = value
	case "GATEWAYPORTS":
		h.GatewayPorts = value
	case "GLOBALKNOWNHOSTSFILE":
		h.GlobalKnownHostsFile = value
	case "GSSAPIAUTHENTICATION":
		h.GSSAPIAuthentication = value
	case "GSSAPIKEYEXCHANGE":
		h.GSSAPIKeyExchange = value
	case "GSSAPICLIENTIDENTITY":
		h.GSSAPIClientIdentity = value
	case "GSSAPIDELEGATECREDENTIALS":
		h.GSSAPIDelegateCredentials = value
	case "GSSAPIRENEWALFORCESREKEY":
		h.GSSAPIRenewalForcesRekey = value
	case "GSSAPITRUSTDNS":
		h.GSSAPITrustDns = value
	case "HASHKNOWNHOSTS":
		h.HashKnownHosts = value
	case "HOSTBASEDAUTHENTICATION":
		h.HostbasedAuthentication = value
	case "HOSTKEYALGORITHMS":
		h.HostKeyAlgorithms = value
	case "HOSTKEYALIAS":
		h.HostKeyAlias = value
	case "HOSTNAME":
		h.HostName = value
	case "IDENTITIESONLY":
		h.IdentitiesOnly = value
	case "IDENTITYFILE":
		h.IdentityFile = value
	case "KBDINTERACTIVEAUTHENTICATION":
		h.KbdInteractiveAuthentication = value
	case "KBDINTERACTIVEDEVICES":
		h.KbdInteractiveDevices = value
	case "LOCALCOMMAND":
		h.LocalCommand = value
	case "LOCALFORWARD":
		h.LocalForward = value
	case "LOGLEVEL":
		h.LogLevel = value
	case "MACS":
		h.MACs = value
	case "NOHOSTAUTHENTICATIONFORLOCALHOST":
		h.NoHostAuthenticationForLocalhost = value
	case "PREFERREDAUTHENTICATIONS":
		h.PreferredAuthentications = value
	case "PROTOCOL":
		h.Protocol = value
	case "PROXYCOMMAND":
		h.ProxyCommand = value
	case "PUBKEYAUTHENTICATION":
		h.PubkeyAuthentication = value
	case "REMOTEFORWARD":
		h.RemoteForward = value
	case "RHOSTSRSAAUTHENTICATION":
		h.RhostsRSAAuthentication = value
	case "RSAAUTHENTICATION":
		h.RSAAuthentication = value
	case "SENDENV":
		h.SendEnv = value
	case "SERVERALIVECOUNTMAX":
		h.ServerAliveCountMax = value
	case "SERVERALIVEINTERVAL":
		h.ServerAliveInterval = value
	case "SMARTCARDDEVICE":
		h.SmartcardDevice = value
	case "STRICTHOSTKEYCHECKING":
		h.StrictHostKeyChecking = value
	case "TCPKEEPALIVE":
		h.TCPKeepAlive = value
	case "TUNNEL":
		h.Tunnel = value
	case "TUNNELDEVICE":
		h.TunnelDevice = value
	case "USEPRIVILEGEDPORT":
		h.UsePrivilegedPort = value
	case "USERKNOWNHOSTSFILE":
		h.UserKnownHostsFile = value
	case "VERIFYHOSTKEYDNS":
		h.VerifyHostKeyDNS = value
	case "VISUALHOSTKEY":
		h.VisualHostKey = value
	}
}

func (h *Host) GetFieldNames() []string {
	var fields []string
	v := reflect.ValueOf(h).Elem()
	for i := 0; i < v.NumField(); i++ {
		fields = append(fields, v.Type().Field(i).Name)
	}
	return fields
}
