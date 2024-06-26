package models

// Conf - web gui config
type Conf struct {
	Host          string
	Port          string
	Theme         string
	Color         string
	DirPath       string
	ConfPath      string
	YamlPath      string
	NodePath      string
	Timeout       int
	HistTrim      int
	InfluxAddr    string
	InfluxToken   string
	InfluxOrg     string
	InfluxBucket  string
	InfluxEnable  bool
	InfluxSkipTLS bool
}

// PortItem - one port
type PortItem struct {
	Name  string
	Port  int
	State bool
	Watch bool
}

// AddrToScan - one addr to scan
type AddrToScan struct {
	Name     string
	Addr     string
	PortMap  map[int]PortItem
	Total    int
	Watching int
	Online   int
	Offline  int
}

// HistState - one history state
type HistState struct {
	Date  string
	State bool
}

// HistData - all data for Addr:Port
type HistData struct {
	Name     string
	Addr     string
	Port     int
	PortName string
	State    []HistState
	NowState bool
}

// GuiData - web gui data
type GuiData struct {
	Config  Conf
	Themes  []string
	Version string
	OneAddr AddrToScan
}
