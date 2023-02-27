package mesopotamia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈尼HaniBarony struct {
	feud.BaseBarony
}

var BHani哈尼 feud.Barony = &哈尼HaniBarony{}

func init() {
    f := BHani哈尼.(*哈尼HaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hani",
		TitleName: "哈尼",
		TitleCode: "b_hani",
	}
}
