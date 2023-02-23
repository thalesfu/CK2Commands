package vorotynsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 索莫沃SomovoBarony struct {
	feud.BaseBarony
}

var BSomovo索莫沃 feud.Barony = &索莫沃SomovoBarony{}

func init() {
	f := BSomovo索莫沃.(*索莫沃SomovoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "somovo",
		TitleName: "索莫沃",
		TitleCode: "b_somovo",
	}
}
