package mansia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 涅夫捷尤甘斯克NefteyuganskBarony struct {
	feud.BaseBarony
}

var BNefteyugansk涅夫捷尤甘斯克 feud.Barony = &涅夫捷尤甘斯克NefteyuganskBarony{}

func init() {
    f := BNefteyugansk涅夫捷尤甘斯克.(*涅夫捷尤甘斯克NefteyuganskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nefteyugansk",
		TitleName: "涅夫捷尤甘斯克",
		TitleCode: "b_nefteyugansk",
	}
}
