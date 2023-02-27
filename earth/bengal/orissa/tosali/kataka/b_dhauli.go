package kataka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陶利DhauliBarony struct {
	feud.BaseBarony
}

var BDhauli陶利 feud.Barony = &陶利DhauliBarony{}

func init() {
    f := BDhauli陶利.(*陶利DhauliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dhauli",
		TitleName: "陶利",
		TitleCode: "b_dhauli",
	}
}
