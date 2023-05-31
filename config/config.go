package config

const Version = "heplify-server 1.59.3"

var Setting HeplifyServer

type HeplifyServer struct {
	HEPAddr            string   `default:"0.0.0.0:9060"`
	HEPTCPAddr         string   `default:""`
	HEPTLSAddr         string   `default:""`
	HEPWSAddr          string   `default:""`
	ESAddr             string   `default:""`
	ESDiscovery        bool     `default:"true"`
	HEPv2Enable        bool     `default:"true"`
	ESUser             string   `default:""`
	ESPass             string   `default:""`
	LokiURL            string   `default:""`
	LokiBulk           int      `default:"400"`
	LokiTimer          int      `default:"4"`
	LokiBuffer         int      `default:"100000"`
	LokiHEPFilter      []int    `default:"1,5,100"`
	ForceHEPPayload    []int    `default:""`
	PromAddr           string   `default:":9096"`
	PromTargetIP       string   `default:""`
	PromTargetName     string   `default:""`
	DBShema            string   `default:"homer5"`
	DBDriver           string   `default:"mysql"`
	DBAddr             string   `default:"localhost:3306"`
	DBSSLMode          string   `default:"disable"`
	DBUser             string   `default:"root"`
	DBPass             string   `default:""`
	DBDataTable        string   `default:"homer_data"`
	DBConfTable        string   `default:"homer_configuration"`
	DBBulk             int      `default:"400"`
	DBTimer            int      `default:"4"`
	DBBuffer           int      `default:"400000"`
	DBWorker           int      `default:"8"`
	DBRotate           bool     `default:"true"`
	DBPartLog          string   `default:"2h"`
	DBPartIsup         string   `default:"6h"`
	DBPartSip          string   `default:"2h"`
	DBPartDiameter     string   `default:"6h"`
	DBPartQos          string   `default:"6h"`
	DBDropDays         int      `default:"14"`
	DBDropDaysCall     int      `default:"0"`
	DBDropDaysRegister int      `default:"0"`
	DBDropDaysDefault  int      `default:"0"`
	DBDropOnStart      bool     `default:"false"`
	DBUsageProtection  bool     `default:"false"`
	DBUsageScheme      string   `default:"percentage"`
	DBPercentageUsage  string   `default:"80%"`
	DBMaxSize          string   `default:"20GB"`
	DBProcDropLimit    int      `default:"2"`
	Dedup              bool     `default:"false"`
	DiscardMethod      []string `default:""`
	CensorMethod       []string `default:""`
	AlegIDs            []string `default:""`
	ForceALegID        bool     `default:"false"`
	CustomHeader       []string `default:""`
	IgnoreCaseCH       bool     `default:"false"`
	SIPHeader          []string `default:"ruri_user,ruri_domain,from_user,from_tag,to_user,callid,cseq,method,user_agent"`
	LogDbg             string   `default:""`
	LogLvl             string   `default:"info"`
	LogStd             bool     `default:"false"`
	LogSys             bool     `default:"false"`
	Config             string   `default:"./heplify-server.toml"`
	ConfigHTTPAddr     string   `default:""`
	ConfigHTTPPW       string   `default:""`
	Version            bool     `default:"false"`
	ScriptEnable       bool     `default:"false"`
	ScriptEngine       string   `default:"lua"`
	ScriptFolder       string   `default:""`
	ScriptHEPFilter    []int    `default:"1,5,100"`
	TLSCertFolder      string   `default:"."`
	TLSMinVersion      string   `default:"1.2"`
}
