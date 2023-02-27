package zanjan_abhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎汗杰ZahanjBarony struct {
	feud.BaseBarony
}

var BZahanj扎汗杰 feud.Barony = &扎汗杰ZahanjBarony{}

func init() {
    f := BZahanj扎汗杰.(*扎汗杰ZahanjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zahanj",
		TitleName: "扎汗杰",
		TitleCode: "b_zahanj",
	}
}
