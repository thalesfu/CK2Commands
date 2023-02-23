package leh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 楚姆唐ChumathangBarony struct {
	feud.BaseBarony
}

var BChumathang楚姆唐 feud.Barony = &楚姆唐ChumathangBarony{}

func init() {
	f := BChumathang楚姆唐.(*楚姆唐ChumathangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chumathang",
		TitleName: "楚姆唐",
		TitleCode: "b_chumathang",
	}
}
