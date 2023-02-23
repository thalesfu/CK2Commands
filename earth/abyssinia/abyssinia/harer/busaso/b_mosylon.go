package busaso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫塞隆MosylonBarony struct {
	feud.BaseBarony
}

var BMosylon莫塞隆 feud.Barony = &莫塞隆MosylonBarony{}

func init() {
	f := BMosylon莫塞隆.(*莫塞隆MosylonBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mosylon",
		TitleName: "莫塞隆",
		TitleCode: "b_mosylon",
	}
}
