package narbonne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯特尔CastresBarony struct {
	feud.BaseBarony
}

var BCastres卡斯特尔 feud.Barony = &卡斯特尔CastresBarony{}

func init() {
	f := BCastres卡斯特尔.(*卡斯特尔CastresBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castres",
		TitleName: "卡斯特尔",
		TitleCode: "b_castres",
	}
}
