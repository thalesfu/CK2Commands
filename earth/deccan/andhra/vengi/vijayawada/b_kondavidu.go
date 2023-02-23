package vijayawada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔达维杜KondaviduBarony struct {
	feud.BaseBarony
}

var BKondavidu孔达维杜 feud.Barony = &孔达维杜KondaviduBarony{}

func init() {
	f := BKondavidu孔达维杜.(*孔达维杜KondaviduBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kondavidu",
		TitleName: "孔达维杜",
		TitleCode: "b_kondavidu",
	}
}
