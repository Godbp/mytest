package struct_tag


type address struct {
	Street     string     `json:"street"`
	Ste        string     `json:"suite,omitempty"`
	City       string     `json:"city"`
	State      string     `json:"state"`
	Zipcode    string     `json:"zipcode"`
	Coordinate *coordinate `json:"coordinate,omitempty"`
}

type coordinate struct {
	Lat float64 `json:"latitude,omitempty"`
	Lng float64 `json:"longitude,omitempty"`
}

func GetCoordinate() *coordinate {
	return nil
}

func GetAddress() *address {
	co := GetCoordinate()
	return &address{
		Street:     "sd",
		Ste:        "1",
		City:       "sc",
		State:      "2",
		Zipcode:    "1",
		Coordinate: co,
	}
}