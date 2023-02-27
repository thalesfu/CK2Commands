package narke

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里瑟贝里亚RisebergaBarony struct {
	feud.BaseBarony
}

var BRiseberga里瑟贝里亚 feud.Barony = &里瑟贝里亚RisebergaBarony{}

func init() {
    f := BRiseberga里瑟贝里亚.(*里瑟贝里亚RisebergaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "riseberga",
		TitleName: "里瑟贝里亚",
		TitleCode: "b_riseberga",
	}
}
