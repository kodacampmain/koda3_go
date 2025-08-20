package intermediate

import "fmt"

type TV struct {
	IsScreenOn bool
}

func (t *TV) On() {
	t.IsScreenOn = true
	fmt.Println("TV sudah menyala")
}
func (t *TV) Off() {
	t.IsScreenOn = false
	fmt.Println("TV sudah mati")
}
func (t *TV) ChangeChannel() {}

type AC struct {
	Temp int
	IsOn bool
}

func (ac *AC) On() {
	ac.Temp = 25
	ac.IsOn = true
	fmt.Println("AC menyala dengan suhu 25")
}
func (ac *AC) Off() {
	ac.IsOn = false
	fmt.Println("AC mati")
}

type ElectricFan struct {
	Speed int
}

func (ef *ElectricFan) On() {
	ef.Speed = 40
	fmt.Println("Kipas menyala dengan kecepatan 40 rpm")
}
func (ef *ElectricFan) Off() {
	ef.Speed = 0
	fmt.Println("Kipas mati")
}
func (ef *ElectricFan) IncreaseSpeed() {}

type Electronic interface {
	On()
	Off()
}

func UniversalRemote(instruction string, electronic Electronic) {
	switch instruction {
	case "ON":
		// menyalakan
		electronic.On()
	case "OFF":
		// matikan
		electronic.Off()
	default:
		return
	}
}
