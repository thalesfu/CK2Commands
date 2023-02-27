package bharauli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 摩尼迦城ManikpurBarony struct {
	feud.BaseBarony
}

var BManikpur摩尼迦城 feud.Barony = &摩尼迦城ManikpurBarony{}

func init() {
    f := BManikpur摩尼迦城.(*摩尼迦城ManikpurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manikpur",
		TitleName: "摩尼迦城",
		TitleCode: "b_manikpur",
	}
}
