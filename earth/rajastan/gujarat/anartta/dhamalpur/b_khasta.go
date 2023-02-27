package dhamalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦多KhastaBarony struct {
	feud.BaseBarony
}

var BKhasta迦多 feud.Barony = &迦多KhastaBarony{}

func init() {
    f := BKhasta迦多.(*迦多KhastaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khasta",
		TitleName: "迦多",
		TitleCode: "b_khasta",
	}
}
