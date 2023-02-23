package almaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿力麻里AlmaliqBarony struct {
	feud.BaseBarony
}

var BAlmaliq阿力麻里 feud.Barony = &阿力麻里AlmaliqBarony{}

func init() {
	f := BAlmaliq阿力麻里.(*阿力麻里AlmaliqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almaliq",
		TitleName: "阿力麻里",
		TitleCode: "b_almaliq",
	}
}
