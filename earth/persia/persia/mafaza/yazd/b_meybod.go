package yazd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梅博德MeybodBarony struct {
	feud.BaseBarony
}

var BMeybod梅博德 feud.Barony = &梅博德MeybodBarony{}

func init() {
    f := BMeybod梅博德.(*梅博德MeybodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "meybod",
		TitleName: "梅博德",
		TitleCode: "b_meybod",
	}
}
