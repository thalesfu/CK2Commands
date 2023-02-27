package murcia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡亚萨城MedinasiyasaBarony struct {
	feud.BaseBarony
}

var BMedinasiyasa锡亚萨城 feud.Barony = &锡亚萨城MedinasiyasaBarony{}

func init() {
    f := BMedinasiyasa锡亚萨城.(*锡亚萨城MedinasiyasaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "medinasiyasa",
		TitleName: "锡亚萨城",
		TitleCode: "b_medinasiyasa",
	}
}
