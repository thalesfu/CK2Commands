package cadiz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 加的斯CadizBarony struct {
	feud.BaseBarony
}

var BCadiz加的斯 feud.Barony = &加的斯CadizBarony{}

func init() {
    f := BCadiz加的斯.(*加的斯CadizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cadiz",
		TitleName: "加的斯",
		TitleCode: "b_cadiz",
	}
}
