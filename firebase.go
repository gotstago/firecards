// firego project main.go
package main

import (
	"fmt"
	"github.com/CloudCom/fireauth"
	"github.com/CloudCom/firego"
)

func main() {
	gen := fireauth.New("0AZYH24U8xgVMTmTDDpnveMZOvEsZruiViFI8S7q")
	fmt.Println("Hello World!", gen)
	data := fireauth.Data{"uid": "1"}
	token, err := gen.CreateToken(data, nil)
	if err != nil {
		//log.Fatal(err)
		println("my error: ", err)
	}
	println("my token: ", token)

	//f := firego.New("https://my-to-do.firebaseio.com/sample")
	games := firego.New("https://firstgame.firebaseio.com/games")
	//f.Auth(token)
	games.Auth(token)
	// fetch the data and hydrate v with the result
	//var v map[string]interface{}
	//	if err := f.Value(&v); err != nil {
	//		//log.Fatal(err)
	//		println("my error: ", err)
	//	}
	//fmt.Printf("%s\n", v)
	// update Firebase with new data
	//newData := map[string]string{"foorrr": "barrrr"}
	//f.Value()
	//	if err := f.Update(newData); err != nil {
	//		//log.Fatal(err)
	//		println("my error: ", err)
	//	}
	// watch for updates
	notifications := make(chan firego.Event)
	//	userEvents := make(chan firego.Event)
	//	f.Watch(notifications)
	//	e := <-notifications
	//	fmt.Printf("type=%s path=%s data=%v", e.Type, e.Path, e.Data)
	if err := games.Watch(notifications); err != nil {
		//		log.Fatal(err)
		println("my error: ", err)
	}

	defer games.StopWatching()
	for e := range notifications {
		//fmt.Printf("Event %#v\n", event)
		fmt.Printf("type=%s path=%s data=%#v\n", e.Type, e.Path, e.Data)
	}
	fmt.Printf("Notifications have stopped")
}
