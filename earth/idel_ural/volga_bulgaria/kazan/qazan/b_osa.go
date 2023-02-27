package qazan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥萨OsaBarony struct {
	feud.BaseBarony
}

var BOsa奥萨 feud.Barony = &奥萨OsaBarony{}

func init() {
    f := BOsa奥萨.(*奥萨OsaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "osa",
		TitleName: "奥萨",
		TitleCode: "b_osa",
	}
}
