package zachlumia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔丘拉KorculaBarony struct {
	feud.BaseBarony
}

var BKorcula科尔丘拉 feud.Barony = &科尔丘拉KorculaBarony{}

func init() {
    f := BKorcula科尔丘拉.(*科尔丘拉KorculaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korcula",
		TitleName: "科尔丘拉",
		TitleCode: "b_korcula",
	}
}
