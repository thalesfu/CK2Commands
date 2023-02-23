package torres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨萨里SassariBarony struct {
	feud.BaseBarony
}

var BSassari萨萨里 feud.Barony = &萨萨里SassariBarony{}

func init() {
	f := BSassari萨萨里.(*萨萨里SassariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sassari",
		TitleName: "萨萨里",
		TitleCode: "b_sassari",
	}
}
