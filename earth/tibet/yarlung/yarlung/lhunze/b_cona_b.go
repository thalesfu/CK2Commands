package lhunze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 错那Cona_bBarony struct {
	feud.BaseBarony
}

var BCona_b错那 feud.Barony = &错那Cona_bBarony{}

func init() {
    f := BCona_b错那.(*错那Cona_bBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cona_b",
		TitleName: "错那",
		TitleCode: "b_cona_b",
	}
}
