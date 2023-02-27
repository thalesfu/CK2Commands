package carrick

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特恩贝里TurnberryBarony struct {
	feud.BaseBarony
}

var BTurnberry特恩贝里 feud.Barony = &特恩贝里TurnberryBarony{}

func init() {
    f := BTurnberry特恩贝里.(*特恩贝里TurnberryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turnberry",
		TitleName: "特恩贝里",
		TitleCode: "b_turnberry",
	}
}
