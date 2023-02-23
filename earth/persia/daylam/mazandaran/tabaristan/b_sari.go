package tabaristan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨里SariBarony struct {
	feud.BaseBarony
}

var BSari萨里 feud.Barony = &萨里SariBarony{}

func init() {
	f := BSari萨里.(*萨里SariBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sari",
		TitleName: "萨里",
		TitleCode: "b_sari",
	}
}
