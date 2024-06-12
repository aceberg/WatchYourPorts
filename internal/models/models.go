package models

// Conf - web gui config
type Conf struct {
	Host     string
	Port     string
	Theme    string
	Color    string
	Icon     string
	DBPath   string
	DirPath  string
	ConfPath string
	NodePath string
}

// PortItem - one port
type PortItem struct {
	Name  string
	Port  int
	State bool
}

// AddrToScan - one addr to scan
type AddrToScan struct {
	Addr     string
	PortList []PortItem
}

// GuiData - web gui data
type GuiData struct {
	Config  Conf
	Themes  []string
	Version string
}
