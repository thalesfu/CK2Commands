package najera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔法拉AlfaraBarony struct {
	feud.BaseBarony
}

var BAlfara阿尔法拉 feud.Barony = &阿尔法拉AlfaraBarony{}

func init() {
    f := BAlfara阿尔法拉.(*阿尔法拉AlfaraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alfara",
		TitleName: "阿尔法拉",
		TitleCode: "b_alfara",
	}
}
