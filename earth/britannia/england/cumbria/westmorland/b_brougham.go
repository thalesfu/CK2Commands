package westmorland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布鲁厄姆BroughamBarony struct {
	feud.BaseBarony
}

var BBrougham布鲁厄姆 feud.Barony = &布鲁厄姆BroughamBarony{}

func init() {
	f := BBrougham布鲁厄姆.(*布鲁厄姆BroughamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brougham",
		TitleName: "布鲁厄姆",
		TitleCode: "b_brougham",
	}
}
