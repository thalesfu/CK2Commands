package istria

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲乌梅FiumeBarony struct {
	feud.BaseBarony
}

var BFiume菲乌梅 feud.Barony = &菲乌梅FiumeBarony{}

func init() {
    f := BFiume菲乌梅.(*菲乌梅FiumeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fiume",
		TitleName: "菲乌梅",
		TitleCode: "b_fiume",
	}
}
