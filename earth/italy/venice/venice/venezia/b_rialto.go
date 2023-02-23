package venezia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 里亚尔托RialtoBarony struct {
	feud.BaseBarony
}

var BRialto里亚尔托 feud.Barony = &里亚尔托RialtoBarony{}

func init() {
	f := BRialto里亚尔托.(*里亚尔托RialtoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rialto",
		TitleName: "里亚尔托",
		TitleCode: "b_rialto",
	}
}
