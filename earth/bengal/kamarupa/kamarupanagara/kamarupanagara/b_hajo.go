package kamarupanagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诃阇乌HajoBarony struct {
	feud.BaseBarony
}

var BHajo诃阇乌 feud.Barony = &诃阇乌HajoBarony{}

func init() {
    f := BHajo诃阇乌.(*诃阇乌HajoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hajo",
		TitleName: "诃阇乌",
		TitleCode: "b_hajo",
	}
}
