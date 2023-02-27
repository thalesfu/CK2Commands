package qwivir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑格萨尔SangsarBarony struct {
	feud.BaseBarony
}

var BSangsar桑格萨尔 feud.Barony = &桑格萨尔SangsarBarony{}

func init() {
    f := BSangsar桑格萨尔.(*桑格萨尔SangsarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangsar",
		TitleName: "桑格萨尔",
		TitleCode: "b_sangsar",
	}
}
