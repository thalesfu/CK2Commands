package moray

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 因弗内斯InvernessBarony struct {
	feud.BaseBarony
}

var BInverness因弗内斯 feud.Barony = &因弗内斯InvernessBarony{}

func init() {
    f := BInverness因弗内斯.(*因弗内斯InvernessBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "inverness",
		TitleName: "因弗内斯",
		TitleCode: "b_inverness",
	}
}
