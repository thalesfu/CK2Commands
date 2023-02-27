package herjedalen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坦奈斯TannasBarony struct {
	feud.BaseBarony
}

var BTannas坦奈斯 feud.Barony = &坦奈斯TannasBarony{}

func init() {
    f := BTannas坦奈斯.(*坦奈斯TannasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tannas",
		TitleName: "坦奈斯",
		TitleCode: "b_tannas",
	}
}
