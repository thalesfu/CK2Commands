package belz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 洛帕京LopatynBarony struct {
	feud.BaseBarony
}

var BLopatyn洛帕京 feud.Barony = &洛帕京LopatynBarony{}

func init() {
    f := BLopatyn洛帕京.(*洛帕京LopatynBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lopatyn",
		TitleName: "洛帕京",
		TitleCode: "b_lopatyn",
	}
}
