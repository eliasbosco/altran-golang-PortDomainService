package types

//PortsBody all the properties related to port
type Ports struct {
	PortId      string
	Name        string
	City        string
	Country     string
	Alias       []string
	Regions     []string
	Coordinates []float32
	Province    string
	Timezone    string
	Unlocs      []string
	Code        string
}
