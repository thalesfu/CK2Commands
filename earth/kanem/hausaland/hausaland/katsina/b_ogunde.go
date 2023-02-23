package katsina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥贡德OgundeBarony struct {
	feud.BaseBarony
}

var BOgunde奥贡德 feud.Barony = &奥贡德OgundeBarony{}

func init() {
	f := BOgunde奥贡德.(*奥贡德OgundeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ogunde",
		TitleName: "奥贡德",
		TitleCode: "b_ogunde",
	}
}
