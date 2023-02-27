package loulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 东华门DonghuamenBarony struct {
	feud.BaseBarony
}

var BDonghuamen东华门 feud.Barony = &东华门DonghuamenBarony{}

func init() {
    f := BDonghuamen东华门.(*东华门DonghuamenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "donghuamen",
		TitleName: "东华门",
		TitleCode: "b_donghuamen",
	}
}
