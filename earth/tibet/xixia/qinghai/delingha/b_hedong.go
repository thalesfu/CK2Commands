package delingha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 河东HedongBarony struct {
	feud.BaseBarony
}

var BHedong河东 feud.Barony = &河东HedongBarony{}

func init() {
    f := BHedong河东.(*河东HedongBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hedong",
		TitleName: "河东",
		TitleCode: "b_hedong",
	}
}
