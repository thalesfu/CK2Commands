package nedong

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 桑日SangriBarony struct {
	feud.BaseBarony
}

var BSangri桑日 feud.Barony = &桑日SangriBarony{}

func init() {
    f := BSangri桑日.(*桑日SangriBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sangri",
		TitleName: "桑日",
		TitleCode: "b_sangri",
	}
}
