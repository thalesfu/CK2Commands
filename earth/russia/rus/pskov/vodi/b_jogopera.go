package vodi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 约格佩雷JogoperaBarony struct {
	feud.BaseBarony
}

var BJogopera约格佩雷 feud.Barony = &约格佩雷JogoperaBarony{}

func init() {
    f := BJogopera约格佩雷.(*约格佩雷JogoperaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jogopera",
		TitleName: "约格佩雷",
		TitleCode: "b_jogopera",
	}
}
