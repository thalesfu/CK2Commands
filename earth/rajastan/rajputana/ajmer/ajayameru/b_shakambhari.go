package ajayameru

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 舍剑婆梨ShakambhariBarony struct {
	feud.BaseBarony
}

var BShakambhari舍剑婆梨 feud.Barony = &舍剑婆梨ShakambhariBarony{}

func init() {
	f := BShakambhari舍剑婆梨.(*舍剑婆梨ShakambhariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shakambhari",
		TitleName: "舍剑婆梨",
		TitleCode: "b_shakambhari",
	}
}
