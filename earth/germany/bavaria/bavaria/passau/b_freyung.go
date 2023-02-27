package passau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 弗赖翁FreyungBarony struct {
	feud.BaseBarony
}

var BFreyung弗赖翁 feud.Barony = &弗赖翁FreyungBarony{}

func init() {
    f := BFreyung弗赖翁.(*弗赖翁FreyungBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "freyung",
		TitleName: "弗赖翁",
		TitleCode: "b_freyung",
	}
}
