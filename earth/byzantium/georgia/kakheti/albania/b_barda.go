package albania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔达BardaBarony struct {
	feud.BaseBarony
}

var BBarda巴尔达 feud.Barony = &巴尔达BardaBarony{}

func init() {
	f := BBarda巴尔达.(*巴尔达BardaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barda",
		TitleName: "巴尔达",
		TitleCode: "b_barda",
	}
}
