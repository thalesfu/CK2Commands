package kakheti

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉米GremiBarony struct {
	feud.BaseBarony
}

var BGremi吉米 feud.Barony = &吉米GremiBarony{}

func init() {
    f := BGremi吉米.(*吉米GremiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gremi",
		TitleName: "吉米",
		TitleCode: "b_gremi",
	}
}
