package trinkitat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 法拉塔FarataBarony struct {
	feud.BaseBarony
}

var BFarata法拉塔 feud.Barony = &法拉塔FarataBarony{}

func init() {
	f := BFarata法拉塔.(*法拉塔FarataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "farata",
		TitleName: "法拉塔",
		TitleCode: "b_farata",
	}
}
