package zanjan_abhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 莫扎法里MozaffariBarony struct {
	feud.BaseBarony
}

var BMozaffari莫扎法里 feud.Barony = &莫扎法里MozaffariBarony{}

func init() {
    f := BMozaffari莫扎法里.(*莫扎法里MozaffariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mozaffari",
		TitleName: "莫扎法里",
		TitleCode: "b_mozaffari",
	}
}
