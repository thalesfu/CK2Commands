package kolva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼济瓦NizvaBarony struct {
	feud.BaseBarony
}

var BNizva尼济瓦 feud.Barony = &尼济瓦NizvaBarony{}

func init() {
    f := BNizva尼济瓦.(*尼济瓦NizvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nizva",
		TitleName: "尼济瓦",
		TitleCode: "b_nizva",
	}
}
