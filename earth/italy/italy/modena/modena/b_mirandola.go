package modena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米兰多拉MirandolaBarony struct {
	feud.BaseBarony
}

var BMirandola米兰多拉 feud.Barony = &米兰多拉MirandolaBarony{}

func init() {
    f := BMirandola米兰多拉.(*米兰多拉MirandolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mirandola",
		TitleName: "米兰多拉",
		TitleCode: "b_mirandola",
	}
}
