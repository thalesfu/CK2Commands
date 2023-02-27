package innsbruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 延巴赫JenbachBarony struct {
	feud.BaseBarony
}

var BJenbach延巴赫 feud.Barony = &延巴赫JenbachBarony{}

func init() {
    f := BJenbach延巴赫.(*延巴赫JenbachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jenbach",
		TitleName: "延巴赫",
		TitleCode: "b_jenbach",
	}
}
