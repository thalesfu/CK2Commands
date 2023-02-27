package kiev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尤里耶夫YurievBarony struct {
	feud.BaseBarony
}

var BYuriev尤里耶夫 feud.Barony = &尤里耶夫YurievBarony{}

func init() {
    f := BYuriev尤里耶夫.(*尤里耶夫YurievBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yuriev",
		TitleName: "尤里耶夫",
		TitleCode: "b_yuriev",
	}
}
