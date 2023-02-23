package desmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉尼BlarneyBarony struct {
	feud.BaseBarony
}

var BBlarney布拉尼 feud.Barony = &布拉尼BlarneyBarony{}

func init() {
	f := BBlarney布拉尼.(*布拉尼BlarneyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blarney",
		TitleName: "布拉尼",
		TitleCode: "b_blarney",
	}
}
