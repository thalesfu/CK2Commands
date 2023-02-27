package tourraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 菲尔布瓦FierboisBarony struct {
	feud.BaseBarony
}

var BFierbois菲尔布瓦 feud.Barony = &菲尔布瓦FierboisBarony{}

func init() {
    f := BFierbois菲尔布瓦.(*菲尔布瓦FierboisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "fierbois",
		TitleName: "菲尔布瓦",
		TitleCode: "b_fierbois",
	}
}
