package vijayawada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 慕那吠罗MonvelBarony struct {
	feud.BaseBarony
}

var BMonvel慕那吠罗 feud.Barony = &慕那吠罗MonvelBarony{}

func init() {
    f := BMonvel慕那吠罗.(*慕那吠罗MonvelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "monvel",
		TitleName: "慕那吠罗",
		TitleCode: "b_monvel",
	}
}
