package rugen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔特BarthBarony struct {
	feud.BaseBarony
}

var BBarth巴尔特 feud.Barony = &巴尔特BarthBarony{}

func init() {
    f := BBarth巴尔特.(*巴尔特BarthBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barth",
		TitleName: "巴尔特",
		TitleCode: "b_barth",
	}
}
