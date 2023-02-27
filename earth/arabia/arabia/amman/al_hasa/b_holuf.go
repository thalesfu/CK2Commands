package al_hasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 胡富夫HolufBarony struct {
	feud.BaseBarony
}

var BHoluf胡富夫 feud.Barony = &胡富夫HolufBarony{}

func init() {
    f := BHoluf胡富夫.(*胡富夫HolufBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "holuf",
		TitleName: "胡富夫",
		TitleCode: "b_holuf",
	}
}
