package julich

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鲁尔蒙德RoermondBarony struct {
	feud.BaseBarony
}

var BRoermond鲁尔蒙德 feud.Barony = &鲁尔蒙德RoermondBarony{}

func init() {
    f := BRoermond鲁尔蒙德.(*鲁尔蒙德RoermondBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "roermond",
		TitleName: "鲁尔蒙德",
		TitleCode: "b_roermond",
	}
}
