package tarragona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雷乌斯ReusBarony struct {
	feud.BaseBarony
}

var BReus雷乌斯 feud.Barony = &雷乌斯ReusBarony{}

func init() {
	f := BReus雷乌斯.(*雷乌斯ReusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reus",
		TitleName: "雷乌斯",
		TitleCode: "b_reus",
	}
}
