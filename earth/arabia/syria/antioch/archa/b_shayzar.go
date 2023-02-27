package archa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西扎ShayzarBarony struct {
	feud.BaseBarony
}

var BShayzar西扎 feud.Barony = &西扎ShayzarBarony{}

func init() {
    f := BShayzar西扎.(*西扎ShayzarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shayzar",
		TitleName: "西扎",
		TitleCode: "b_shayzar",
	}
}
