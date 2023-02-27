package hellas

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 柳城IteaBarony struct {
	feud.BaseBarony
}

var BItea柳城 feud.Barony = &柳城IteaBarony{}

func init() {
    f := BItea柳城.(*柳城IteaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "itea",
		TitleName: "柳城",
		TitleCode: "b_itea",
	}
}
