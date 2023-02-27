package napata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯里迈KarimaBarony struct {
	feud.BaseBarony
}

var BKarima凯里迈 feud.Barony = &凯里迈KarimaBarony{}

func init() {
    f := BKarima凯里迈.(*凯里迈KarimaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karima",
		TitleName: "凯里迈",
		TitleCode: "b_karima",
	}
}
