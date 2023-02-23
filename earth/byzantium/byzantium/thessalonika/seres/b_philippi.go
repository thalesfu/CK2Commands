package seres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 腓立比PhilippiBarony struct {
	feud.BaseBarony
}

var BPhilippi腓立比 feud.Barony = &腓立比PhilippiBarony{}

func init() {
	f := BPhilippi腓立比.(*腓立比PhilippiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "philippi",
		TitleName: "腓立比",
		TitleCode: "b_philippi",
	}
}
