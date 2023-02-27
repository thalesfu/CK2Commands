package durham

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷比RabyBarony struct {
	feud.BaseBarony
}

var BRaby雷比 feud.Barony = &雷比RabyBarony{}

func init() {
    f := BRaby雷比.(*雷比RabyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raby",
		TitleName: "雷比",
		TitleCode: "b_raby",
	}
}
