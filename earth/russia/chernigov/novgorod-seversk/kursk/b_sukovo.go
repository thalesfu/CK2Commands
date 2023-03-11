package kursk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 苏科沃SukovoBarony struct {
	feud.BaseBarony
}

var BSukovo苏科沃 feud.Barony = &苏科沃SukovoBarony{}

func init() {
    f := BSukovo苏科沃.(*苏科沃SukovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sukovo",
		TitleName: "苏科沃",
		TitleCode: "b_sukovo",
	}
}
