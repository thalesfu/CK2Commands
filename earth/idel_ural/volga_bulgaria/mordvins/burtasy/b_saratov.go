package burtasy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉托夫SaratovBarony struct {
	feud.BaseBarony
}

var BSaratov萨拉托夫 feud.Barony = &萨拉托夫SaratovBarony{}

func init() {
    f := BSaratov萨拉托夫.(*萨拉托夫SaratovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saratov",
		TitleName: "萨拉托夫",
		TitleCode: "b_saratov",
	}
}
