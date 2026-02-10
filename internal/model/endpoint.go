package model

type Endpoint struct {
	IPAddress         string  `json:"ipAddress"`
	StatusMessage     string  `json:"statusMessage"`
	Grade             string  `json:"grade"`
	GradeTrustIgnored string  `json:"gradeTrustIgnored"`
	HasWarnings       bool    `json:"hasWarnings"`
	IsExceptional     bool    `json:"isExceptional"`
	Progress          int     `json:"progress"`
	Duration          int     `json:"duration"`
	ETA               int     `json:"eta"`
	Delegation        int     `json:"delegation"`
	Details           Details `json:"details"`
}

type Details struct {
	HostStartTime       int64         `json:"hostStartTime"`
	Key                 Key           `json:"key"`
	Cert                Cert          `json:"cert"`
	Chain               Chain         `json:"chain"`
	Protocols           []Protocols   `json:"protocols"`
	Suites              Suites        `json:"suites"`
	ServerSignature     string        `json:"serverSignature"`
	PrefixDelegation    bool          `json:"prefixDelegation"`
	NonPrefixDelegation bool          `json:"nonPrefixDelegation"`
	VulnBeast           bool          `json:"vulnBeast"`
	RenegSupport        int           `json:"renegSupport"`
	StsStatus           string        `json:"stsStatus"`
	StsResponseHeader   string        `json:"stsResponseHeader"`
	StsMaxAge           int           `json:"stsMaxAge"`
	StsSubdomains       bool          `json:"stsSubdomains"`
	StsPreload          bool          `json:"stsPreload"`
	SessionResumption   int           `json:"sessionResumption"`
	CompressionMethods  int           `json:"compressionMethods"`
	SupportsNpn         bool          `json:"supportsNpn"`
	SupportsAlpn        bool          `json:"supportsAlpn"`
	SessionTickets      int           `json:"sessionTickets"`
	OcspStapling        bool          `json:"ocspStapling"`
	SniRequired         bool          `json:"sniRequired"`
	HTTPStatusCode      int           `json:"httpStatusCode"`
	SupportsRc4         bool          `json:"supportsRc4"`
	Rc4WithModern       bool          `json:"rc4WithModern"`
	Rc4Only             bool          `json:"rc4Only"`
	ForwardSecrecy      int           `json:"forwardSecrecy"`
	ProtocolIntolerance int           `json:"protocolIntolerance"`
	MiscIntolerance     int           `json:"miscIntolerance"`
	Sims                Sims          `json:"sims"`
	Heartbleed          bool          `json:"heartbleed"`
	Heartbeat           bool          `json:"heartbeat"`
	OpenSslCcs          int           `json:"openSslCcs"`
	OpenSSLLuckyMinus20 int           `json:"openSSLLuckyMinus20"`
	Poodle              bool          `json:"poodle"`
	PoodleTls           int           `json:"poodleTls"`
	FallbackScsv        bool          `json:"fallbackScsv"`
	Freak               bool          `json:"freak"`
	HasSct              int           `json:"hasSct"`
	DhPrimes            []string      `json:"dhPrimes"`
	DhUsesKnownPrimes   int           `json:"dhUsesKnownPrimes"`
	DhYsReuse           bool          `json:"dhYsReuse"`
	Logjam              bool          `json:"logjam"`
	ChaCha20Preference  bool          `json:"chaCha20Preference"`
	HstsPolicy          HstsPolicy    `json:"hstsPolicy"`
	HstsPreloads        []HstsPreload `json:"hstsPreloads"`
	HpkpPolicy          HpkpPolicy    `json:"hpkpPolicy"`
	HpkpRoPolicy        HpkpPolicy    `json:"hpkpRoPolicy"`
}

type Key struct {
	Size       int    `json:"size"`
	Alg        string `json:"alg"`
	DebianFlaw bool   `json:"debianFlaw"`
	Strength   int    `json:"strength"`
}

type Cert struct {
	Subject              string   `json:"subject"`
	CommonNames          []string `json:"commonNames"`
	AltNames             []string `json:"altNames"`
	NotBefore            int64    `json:"notBefore"`
	NotAfter             int64    `json:"notAfter"`
	IssuerSubject        string   `json:"issuerSubject"`
	IssuerLabel          string   `json:"issuerLabel"`
	SigAlg               string   `json:"sigAlg"`
	RevocationInfo       int      `json:"revocationInfo"`
	CrlURIs              []string `json:"crlURIs"`
	OcspURIs             []string `json:"ocspURIs"`
	RevocationStatus     int      `json:"revocationStatus"`
	CrlRevocationStatus  int      `json:"crlRevocationStatus"`
	OcspRevocationStatus int      `json:"ocspRevocationStatus"`
	Sgc                  int      `json:"sgc"`
	Issues               int      `json:"issues"`
	Sct                  bool     `json:"sct"`
	MustStaple           int      `json:"mustStaple"`
	Sha1Hash             string   `json:"sha1Hash"`
	PinSha256            string   `json:"pinSha256"`
}

type Chain struct {
	Certs  []ChainCert `json:"certs"`
	Issues int         `json:"issues"`
}

type ChainCert struct {
	Subject              string `json:"subject"`
	Label                string `json:"label"`
	NotBefore            int64  `json:"notBefore"`
	NotAfter             int64  `json:"notAfter"`
	IssuerSubject        string `json:"issuerSubject"`
	IssuerLabel          string `json:"issuerLabel"`
	SigAlg               string `json:"sigAlg"`
	Issues               int    `json:"issues"`
	KeyAlg               string `json:"keyAlg"`
	KeySize              int    `json:"keySize"`
	KeyStrength          int    `json:"keyStrength"`
	RevocationStatus     int    `json:"revocationStatus"`
	CrlRevocationStatus  int    `json:"crlRevocationStatus"`
	OcspRevocationStatus int    `json:"ocspRevocationStatus"`
	Sha1Hash             string `json:"sha1Hash"`
	PinSha256            string `json:"pinSha256"`
	Raw                  string `json:"raw"`
}

type Protocols struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Suites struct {
	List       []Suite `json:"list"`
	Preference bool    `json:"preference"`
}

type Suite struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	CipherStrength int    `json:"cipherStrength"`
	EcdhBits       int    `json:"ecdhBits,omitempty"`
	EcdhStrength   int    `json:"ecdhStrength,omitempty"`
	DhStrength     int    `json:"dhStrength,omitempty"`
	DhP            int    `json:"dhP,omitempty"`
	DhG            int    `json:"dhG,omitempty"`
	DhYs           int    `json:"dhYs,omitempty"`
}

type Sims struct {
	Results []SimResult `json:"results"`
}

type SimResult struct {
	Client     map[string]any `json:"client"`
	ErrorCode  int            `json:"errorCode"`
	Attempts   int            `json:"attempts"`
	ProtocolId int            `json:"protocolId,omitempty"`
	SuiteId    int            `json:"suiteId,omitempty"`
	KxInfo     string         `json:"kxInfo,omitempty"`
}

type HstsPolicy struct {
	LongMaxAge int               `json:"LONG_MAX_AGE"`
	Header     string            `json:"header"`
	Status     string            `json:"status"`
	MaxAge     int               `json:"maxAge"`
	Directives map[string]string `json:"directives"`
}

type HstsPreload struct {
	Source     string `json:"source"`
	Hostname   string `json:"hostname"`
	Status     string `json:"status"`
	SourceTime int64  `json:"sourceTime"`
}

type HpkpPolicy struct {
	Status      string        `json:"status"`
	Pins        []string      `json:"pins"`
	MatchedPins []string      `json:"matchedPins"`
	Directives  []interface{} `json:"directives"`
}
