package syrt

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥辛基OsinkiBarony struct {
	feud.BaseBarony
}

var BOsinki奥辛基 feud.Barony = &奥辛基OsinkiBarony{}

func init() {
    f := BOsinki奥辛基.(*奥辛基OsinkiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osinki",
		TitleName: "奥辛基",
		TitleCode: "b_osinki",
	}
}
