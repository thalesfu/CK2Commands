package mandapika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 婆伐那迦BawangajaBarony struct {
	feud.BaseBarony
}

var BBawangaja婆伐那迦 feud.Barony = &婆伐那迦BawangajaBarony{}

func init() {
	f := BBawangaja婆伐那迦.(*婆伐那迦BawangajaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bawangaja",
		TitleName: "婆伐那迦",
		TitleCode: "b_bawangaja",
	}
}
