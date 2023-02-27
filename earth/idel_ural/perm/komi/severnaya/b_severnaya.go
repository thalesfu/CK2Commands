package severnaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢韦尔纳亚SevernayaBarony struct {
	feud.BaseBarony
}

var BSevernaya谢韦尔纳亚 feud.Barony = &谢韦尔纳亚SevernayaBarony{}

func init() {
    f := BSevernaya谢韦尔纳亚.(*谢韦尔纳亚SevernayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "severnaya",
		TitleName: "谢韦尔纳亚",
		TitleCode: "b_severnaya",
	}
}
