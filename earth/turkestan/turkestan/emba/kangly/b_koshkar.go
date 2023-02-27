package kangly

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科什卡尔KoshkarBarony struct {
	feud.BaseBarony
}

var BKoshkar科什卡尔 feud.Barony = &科什卡尔KoshkarBarony{}

func init() {
    f := BKoshkar科什卡尔.(*科什卡尔KoshkarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koshkar",
		TitleName: "科什卡尔",
		TitleCode: "b_koshkar",
	}
}
