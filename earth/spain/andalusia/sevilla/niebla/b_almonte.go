package niebla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔蒙特AlmonteBarony struct {
	feud.BaseBarony
}

var BAlmonte阿尔蒙特 feud.Barony = &阿尔蒙特AlmonteBarony{}

func init() {
	f := BAlmonte阿尔蒙特.(*阿尔蒙特AlmonteBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "almonte",
		TitleName: "阿尔蒙特",
		TitleCode: "b_almonte",
	}
}
