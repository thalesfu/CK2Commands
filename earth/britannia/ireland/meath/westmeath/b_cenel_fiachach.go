package westmeath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基内尔菲亚哈赫Cenel_fiachachBarony struct {
	feud.BaseBarony
}

var BCenel_fiachach基内尔菲亚哈赫 feud.Barony = &基内尔菲亚哈赫Cenel_fiachachBarony{}

func init() {
    f := BCenel_fiachach基内尔菲亚哈赫.(*基内尔菲亚哈赫Cenel_fiachachBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cenel_fiachach",
		TitleName: "基内尔菲亚哈赫",
		TitleCode: "b_cenel_fiachach",
	}
}
