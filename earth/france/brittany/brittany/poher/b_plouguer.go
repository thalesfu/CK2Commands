package poher

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普卢盖尔PlouguerBarony struct {
	feud.BaseBarony
}

var BPlouguer普卢盖尔 feud.Barony = &普卢盖尔PlouguerBarony{}

func init() {
	f := BPlouguer普卢盖尔.(*普卢盖尔PlouguerBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "plouguer",
		TitleName: "普卢盖尔",
		TitleCode: "b_plouguer",
	}
}
