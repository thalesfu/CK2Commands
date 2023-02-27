package bornholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥基克比AakirkebyBarony struct {
	feud.BaseBarony
}

var BAakirkeby奥基克比 feud.Barony = &奥基克比AakirkebyBarony{}

func init() {
    f := BAakirkeby奥基克比.(*奥基克比AakirkebyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aakirkeby",
		TitleName: "奥基克比",
		TitleCode: "b_aakirkeby",
	}
}
