package firenze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲耶索莱FiesoleBarony struct {
	feud.BaseBarony
}

var BFiesole菲耶索莱 feud.Barony = &菲耶索莱FiesoleBarony{}

func init() {
    f := BFiesole菲耶索莱.(*菲耶索莱FiesoleBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fiesole",
		TitleName: "菲耶索莱",
		TitleCode: "b_fiesole",
	}
}
