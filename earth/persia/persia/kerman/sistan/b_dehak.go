package sistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 代哈克DehakBarony struct {
	feud.BaseBarony
}

var BDehak代哈克 feud.Barony = &代哈克DehakBarony{}

func init() {
    f := BDehak代哈克.(*代哈克DehakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dehak",
		TitleName: "代哈克",
		TitleCode: "b_dehak",
	}
}
