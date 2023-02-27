package opole

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉齐布日RaciborzBarony struct {
	feud.BaseBarony
}

var BRaciborz拉齐布日 feud.Barony = &拉齐布日RaciborzBarony{}

func init() {
    f := BRaciborz拉齐布日.(*拉齐布日RaciborzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "raciborz",
		TitleName: "拉齐布日",
		TitleCode: "b_raciborz",
	}
}
