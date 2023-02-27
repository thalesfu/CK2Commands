package mahoyadapuram

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡奢迦利KunjakariBarony struct {
	feud.BaseBarony
}

var BKunjakari贡奢迦利 feud.Barony = &贡奢迦利KunjakariBarony{}

func init() {
    f := BKunjakari贡奢迦利.(*贡奢迦利KunjakariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kunjakari",
		TitleName: "贡奢迦利",
		TitleCode: "b_kunjakari",
	}
}
