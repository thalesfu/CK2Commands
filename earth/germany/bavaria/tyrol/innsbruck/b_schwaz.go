package innsbruck

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施瓦茨SchwazBarony struct {
	feud.BaseBarony
}

var BSchwaz施瓦茨 feud.Barony = &施瓦茨SchwazBarony{}

func init() {
    f := BSchwaz施瓦茨.(*施瓦茨SchwazBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schwaz",
		TitleName: "施瓦茨",
		TitleCode: "b_schwaz",
	}
}
