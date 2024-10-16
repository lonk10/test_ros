package aburos

type VehicleType interface {
	Init() error                                       // vehicle initialization
	Start() error                                      // start motor
	SetMode(mode string) error                         // set mavlink
	Move(x float64, y float64, z float64) error        // moves vehicle by x,y,z meters
	GetPosition() error                                // returns gps coordinates of vehicle
	SetPosition(x float64, y float64, z float64) error // moves vehicles a gps coordinates x,y,z
	GetReservedKeywords() []string                     // returns reserved keywords for the vehicle
	GetAdditionalFuncs() map[string]func()             // returns additional functions for the vehicle (eg takeoff function for copter)
}
