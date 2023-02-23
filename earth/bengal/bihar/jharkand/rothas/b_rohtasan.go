package rothas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢诃多珊RohtasanBarony struct {
	feud.BaseBarony
}

var BRohtasan卢诃多珊 feud.Barony = &卢诃多珊RohtasanBarony{}

func init() {
	f := BRohtasan卢诃多珊.(*卢诃多珊RohtasanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rohtasan",
		TitleName: "卢诃多珊",
		TitleCode: "b_rohtasan",
	}
}
