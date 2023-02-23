package daylam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏丹尼耶SoltaniyehBarony struct {
	feud.BaseBarony
}

var BSoltaniyeh苏丹尼耶 feud.Barony = &苏丹尼耶SoltaniyehBarony{}

func init() {
	f := BSoltaniyeh苏丹尼耶.(*苏丹尼耶SoltaniyehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "soltaniyeh",
		TitleName: "苏丹尼耶",
		TitleCode: "b_soltaniyeh",
	}
}
