package somnath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦拉沃尔VeravalBarony struct {
	feud.BaseBarony
}

var BVeraval韦拉沃尔 feud.Barony = &韦拉沃尔VeravalBarony{}

func init() {
    f := BVeraval韦拉沃尔.(*韦拉沃尔VeravalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "veraval",
		TitleName: "韦拉沃尔",
		TitleCode: "b_veraval",
	}
}
