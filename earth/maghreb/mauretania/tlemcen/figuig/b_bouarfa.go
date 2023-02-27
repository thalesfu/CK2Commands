package figuig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布阿尔费BouarfaBarony struct {
	feud.BaseBarony
}

var BBouarfa布阿尔费 feud.Barony = &布阿尔费BouarfaBarony{}

func init() {
    f := BBouarfa布阿尔费.(*布阿尔费BouarfaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bouarfa",
		TitleName: "布阿尔费",
		TitleCode: "b_bouarfa",
	}
}
