package celle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑尔福德HerfordBarony struct {
	feud.BaseBarony
}

var BHerford黑尔福德 feud.Barony = &黑尔福德HerfordBarony{}

func init() {
    f := BHerford黑尔福德.(*黑尔福德HerfordBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "herford",
		TitleName: "黑尔福德",
		TitleCode: "b_herford",
	}
}
