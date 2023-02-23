package khijjingakota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伯赫尔达BahaldaBarony struct {
	feud.BaseBarony
}

var BBahalda伯赫尔达 feud.Barony = &伯赫尔达BahaldaBarony{}

func init() {
	f := BBahalda伯赫尔达.(*伯赫尔达BahaldaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bahalda",
		TitleName: "伯赫尔达",
		TitleCode: "b_bahalda",
	}
}
