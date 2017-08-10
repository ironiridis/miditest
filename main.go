package main

import "github.com/rakyll/portmidi"
// import "github.com/bradfitz/iter"
import "fmt"

func main() {
	var err error
	portmidi.Initialize()
	defer portmidi.Terminate()

	//n := portmidi.CountDevices()
	//for f := range iter.N(n) {
	//	fmt.Printf("%d: %+v\n", f, portmidi.Info(portmidi.DeviceID(f)));
	//}
	
	in, err := portmidi.NewInputStream(3, 1024)
	if err != nil {
		panic(err)
	}
	defer in.Close()
	
	out, err := portmidi.NewOutputStream(2, 1024, 100)
	if err != nil {
		panic(err)
	}
	defer out.Close()
	
MainLoop:
	for {
		evs, err := in.Read(1024)
		if err != nil {
			continue
		}
		for e := range evs {
			fmt.Printf("evs[%d]=%+v\n", e, evs[e]);
			if evs[e].Status == 176 && evs[e].Data1 == 92 {
				break MainLoop
			}
			if evs[e].Status == 176 && evs[e].Data1 >= 81 && evs[e].Data1 <= 88 {
				out.WriteShort(0xB0, 81+(88-evs[e].Data1), 127-evs[e].Data2)
			}
		}
	}
	fmt.Println("Exiting.")
}
