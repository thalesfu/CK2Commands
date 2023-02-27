package khilok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比柳泰BilyutayBarony struct {
	feud.BaseBarony
}

var BBilyutay比柳泰 feud.Barony = &比柳泰BilyutayBarony{}

func init() {
    f := BBilyutay比柳泰.(*比柳泰BilyutayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilyutay",
		TitleName: "比柳泰",
		TitleCode: "b_bilyutay",
	}
}
