package halsingland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺腊拉NorralaBarony struct {
	feud.BaseBarony
}

var BNorrala诺腊拉 feud.Barony = &诺腊拉NorralaBarony{}

func init() {
    f := BNorrala诺腊拉.(*诺腊拉NorralaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "norrala",
		TitleName: "诺腊拉",
		TitleCode: "b_norrala",
	}
}
