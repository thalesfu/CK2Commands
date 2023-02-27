package atholl

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安施拉克斯AnthrachsBarony struct {
	feud.BaseBarony
}

var BAnthrachs安施拉克斯 feud.Barony = &安施拉克斯AnthrachsBarony{}

func init() {
    f := BAnthrachs安施拉克斯.(*安施拉克斯AnthrachsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anthrachs",
		TitleName: "安施拉克斯",
		TitleCode: "b_anthrachs",
	}
}
