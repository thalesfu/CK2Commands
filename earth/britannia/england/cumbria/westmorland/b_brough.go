package westmorland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉夫BroughBarony struct {
	feud.BaseBarony
}

var BBrough布拉夫 feud.Barony = &布拉夫BroughBarony{}

func init() {
	f := BBrough布拉夫.(*布拉夫BroughBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "brough",
		TitleName: "布拉夫",
		TitleCode: "b_brough",
	}
}
