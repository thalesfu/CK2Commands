package al_jawf

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图韦尔TuwayrBarony struct {
	feud.BaseBarony
}

var BTuwayr图韦尔 feud.Barony = &图韦尔TuwayrBarony{}

func init() {
    f := BTuwayr图韦尔.(*图韦尔TuwayrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tuwayr",
		TitleName: "图韦尔",
		TitleCode: "b_tuwayr",
	}
}
