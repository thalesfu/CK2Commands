package candhoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 旃度婆CandhobaBarony struct {
	feud.BaseBarony
}

var BCandhoba旃度婆 feud.Barony = &旃度婆CandhobaBarony{}

func init() {
    f := BCandhoba旃度婆.(*旃度婆CandhobaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "candhoba",
		TitleName: "旃度婆",
		TitleCode: "b_candhoba",
	}
}
