package acre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马纳瓦图ManawatBarony struct {
	feud.BaseBarony
}

var BManawat马纳瓦图 feud.Barony = &马纳瓦图ManawatBarony{}

func init() {
	f := BManawat马纳瓦图.(*马纳瓦图ManawatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manawat",
		TitleName: "马纳瓦图",
		TitleCode: "b_manawat",
	}
}
