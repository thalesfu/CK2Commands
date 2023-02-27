package buhairya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拜哈里耶BuhairyaBarony struct {
	feud.BaseBarony
}

var BBuhairya拜哈里耶 feud.Barony = &拜哈里耶BuhairyaBarony{}

func init() {
    f := BBuhairya拜哈里耶.(*拜哈里耶BuhairyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buhairya",
		TitleName: "拜哈里耶",
		TitleCode: "b_buhairya",
	}
}
