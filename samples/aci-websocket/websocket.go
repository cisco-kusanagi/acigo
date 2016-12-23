package main

import (
	"fmt"
	"github.com/udhos/acigo/aci"
)

func main() {

	a, errNew := aci.New(aci.ClientOptions{Debug: true})
	if errNew != nil {
		fmt.Printf("login new client error: %v\n", errNew)
		return
	}

	// Since credentials have not been specified explicitly under ClientOptions,
	// Login() will use env vars: APIC_HOSTS=host, APIC_USER=username, APIC_PASS=pwd
	errLogin := a.Login()
	if errLogin != nil {
		fmt.Printf("login error: %v\n", errLogin)
		return
	}

	if errSock := a.WebsocketOpen(); errSock != nil {
		fmt.Printf("websocket error: %v\n", errSock)
		return
	}

	subscriptionId, errSub := a.TenantSubscribe()
	if errSub != nil {
		fmt.Printf("tenant subscribe error: %v\n", errSub)
		return
	}

	errAdd := a.TenantAdd("tenant-example", "")
	if errAdd != nil {
		fmt.Printf("tenant add error: %v\n", errAdd)
		return
	}

	errDel := a.TenantDel("tenant-example")
	if errDel != nil {
		fmt.Printf("tenant del error: %v\n", errDel)
		return
	}

	var msg interface{}
	if errRead := a.WebsocketReadJson(&msg); errRead != nil {
		fmt.Printf("websocket read error: %v\n", errRead)
		return
	}

	fmt.Printf("websocket message: %v\n", msg)

	errSubRefresh := a.TenantSubscriptionRefresh(subscriptionId)
	if errSubRefresh != nil {
		fmt.Printf("tenant subscription refresh error: %v", errSubRefresh)
		return
	}

	errLogout := a.Logout()
	if errLogout != nil {
		fmt.Printf("logout error: %v\n", errLogout)
		return
	}
}
