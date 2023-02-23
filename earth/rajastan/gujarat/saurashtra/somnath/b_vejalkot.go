package somnath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠阇罗拘吒VejalkotBarony struct {
	feud.BaseBarony
}

var BVejalkot吠阇罗拘吒 feud.Barony = &吠阇罗拘吒VejalkotBarony{}

func init() {
	f := BVejalkot吠阇罗拘吒.(*吠阇罗拘吒VejalkotBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vejalkot",
		TitleName: "吠阇罗拘吒",
		TitleCode: "b_vejalkot",
	}
}
