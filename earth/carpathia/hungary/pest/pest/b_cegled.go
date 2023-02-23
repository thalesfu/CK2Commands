package pest

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 采格莱德CegledBarony struct {
	feud.BaseBarony
}

var BCegled采格莱德 feud.Barony = &采格莱德CegledBarony{}

func init() {
	f := BCegled采格莱德.(*采格莱德CegledBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cegled",
		TitleName: "采格莱德",
		TitleCode: "b_cegled",
	}
}
