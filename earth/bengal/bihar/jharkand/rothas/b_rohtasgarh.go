package rothas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢诃多娑姞利呬RohtasgarhBarony struct {
	feud.BaseBarony
}

var BRohtasgarh卢诃多娑姞利呬 feud.Barony = &卢诃多娑姞利呬RohtasgarhBarony{}

func init() {
    f := BRohtasgarh卢诃多娑姞利呬.(*卢诃多娑姞利呬RohtasgarhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rohtasgarh",
		TitleName: "卢诃多娑姞利呬",
		TitleCode: "b_rohtasgarh",
	}
}
