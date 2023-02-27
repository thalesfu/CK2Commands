package tiberias

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯蒙特CaymontBarony struct {
	feud.BaseBarony
}

var BCaymont凯蒙特 feud.Barony = &凯蒙特CaymontBarony{}

func init() {
    f := BCaymont凯蒙特.(*凯蒙特CaymontBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "caymont",
		TitleName: "凯蒙特",
		TitleCode: "b_caymont",
	}
}
