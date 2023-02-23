package tinmallal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 廷马拉尔TinmallalBarony struct {
	feud.BaseBarony
}

var BTinmallal廷马拉尔 feud.Barony = &廷马拉尔TinmallalBarony{}

func init() {
	f := BTinmallal廷马拉尔.(*廷马拉尔TinmallalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tinmallal",
		TitleName: "廷马拉尔",
		TitleCode: "b_tinmallal",
	}
}
