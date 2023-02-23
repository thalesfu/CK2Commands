package gyantse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡麦KarmaiBarony struct {
	feud.BaseBarony
}

var BKarmai卡麦 feud.Barony = &卡麦KarmaiBarony{}

func init() {
	f := BKarmai卡麦.(*卡麦KarmaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karmai",
		TitleName: "卡麦",
		TitleCode: "b_karmai",
	}
}
