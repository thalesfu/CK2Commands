package oppland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 福旺FavangBarony struct {
	feud.BaseBarony
}

var BFavang福旺 feud.Barony = &福旺FavangBarony{}

func init() {
    f := BFavang福旺.(*福旺FavangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "favang",
		TitleName: "福旺",
		TitleCode: "b_favang",
	}
}
