package esna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔曼特ArmantBarony struct {
	feud.BaseBarony
}

var BArmant阿尔曼特 feud.Barony = &阿尔曼特ArmantBarony{}

func init() {
    f := BArmant阿尔曼特.(*阿尔曼特ArmantBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "armant",
		TitleName: "阿尔曼特",
		TitleCode: "b_armant",
	}
}
