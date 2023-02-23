package makran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基兹KizBarony struct {
	feud.BaseBarony
}

var BKiz基兹 feud.Barony = &基兹KizBarony{}

func init() {
	f := BKiz基兹.(*基兹KizBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kiz",
		TitleName: "基兹",
		TitleCode: "b_kiz",
	}
}
