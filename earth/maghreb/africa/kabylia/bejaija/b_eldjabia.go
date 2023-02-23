package bejaija

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 德贾比亚EldjabiaBarony struct {
	feud.BaseBarony
}

var BEldjabia德贾比亚 feud.Barony = &德贾比亚EldjabiaBarony{}

func init() {
	f := BEldjabia德贾比亚.(*德贾比亚EldjabiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eldjabia",
		TitleName: "德贾比亚",
		TitleCode: "b_eldjabia",
	}
}
