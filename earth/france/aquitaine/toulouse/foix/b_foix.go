package foix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 富瓦FoixBarony struct {
	feud.BaseBarony
}

var BFoix富瓦 feud.Barony = &富瓦FoixBarony{}

func init() {
	f := BFoix富瓦.(*富瓦FoixBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "foix",
		TitleName: "富瓦",
		TitleCode: "b_foix",
	}
}
