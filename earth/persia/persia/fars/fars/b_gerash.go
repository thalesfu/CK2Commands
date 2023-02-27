package fars

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 盖拉什GerashBarony struct {
	feud.BaseBarony
}

var BGerash盖拉什 feud.Barony = &盖拉什GerashBarony{}

func init() {
    f := BGerash盖拉什.(*盖拉什GerashBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gerash",
		TitleName: "盖拉什",
		TitleCode: "b_gerash",
	}
}
