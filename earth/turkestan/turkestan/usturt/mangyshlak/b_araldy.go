package mangyshlak

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿拉尔德AraldyBarony struct {
	feud.BaseBarony
}

var BAraldy阿拉尔德 feud.Barony = &阿拉尔德AraldyBarony{}

func init() {
    f := BAraldy阿拉尔德.(*阿拉尔德AraldyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "araldy",
		TitleName: "阿拉尔德",
		TitleCode: "b_araldy",
	}
}
