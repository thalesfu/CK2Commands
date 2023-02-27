package danzig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施韦茨SchwetzBarony struct {
	feud.BaseBarony
}

var BSchwetz施韦茨 feud.Barony = &施韦茨SchwetzBarony{}

func init() {
    f := BSchwetz施韦茨.(*施韦茨SchwetzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "schwetz",
		TitleName: "施韦茨",
		TitleCode: "b_schwetz",
	}
}
