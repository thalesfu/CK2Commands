package khijjingakota

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴里帕达BaripadaBarony struct {
	feud.BaseBarony
}

var BBaripada巴里帕达 feud.Barony = &巴里帕达BaripadaBarony{}

func init() {
    f := BBaripada巴里帕达.(*巴里帕达BaripadaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baripada",
		TitleName: "巴里帕达",
		TitleCode: "b_baripada",
	}
}
