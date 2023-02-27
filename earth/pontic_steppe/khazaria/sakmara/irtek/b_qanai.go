package irtek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈奈QanaiBarony struct {
	feud.BaseBarony
}

var BQanai哈奈 feud.Barony = &哈奈QanaiBarony{}

func init() {
    f := BQanai哈奈.(*哈奈QanaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qanai",
		TitleName: "哈奈",
		TitleCode: "b_qanai",
	}
}
