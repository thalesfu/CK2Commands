package pongau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蓬高PongauBarony struct {
	feud.BaseBarony
}

var BPongau蓬高 feud.Barony = &蓬高PongauBarony{}

func init() {
	f := BPongau蓬高.(*蓬高PongauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pongau",
		TitleName: "蓬高",
		TitleCode: "b_pongau",
	}
}
