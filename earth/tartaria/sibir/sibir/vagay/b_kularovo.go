package vagay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库拉罗沃KularovoBarony struct {
	feud.BaseBarony
}

var BKularovo库拉罗沃 feud.Barony = &库拉罗沃KularovoBarony{}

func init() {
    f := BKularovo库拉罗沃.(*库拉罗沃KularovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kularovo",
		TitleName: "库拉罗沃",
		TitleCode: "b_kularovo",
	}
}
