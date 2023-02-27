package sandomierskie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑多梅日SandomierzBarony struct {
	feud.BaseBarony
}

var BSandomierz桑多梅日 feud.Barony = &桑多梅日SandomierzBarony{}

func init() {
    f := BSandomierz桑多梅日.(*桑多梅日SandomierzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sandomierz",
		TitleName: "桑多梅日",
		TitleCode: "b_sandomierz",
	}
}
