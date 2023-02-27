package pavia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博比奥BobbioBarony struct {
	feud.BaseBarony
}

var BBobbio博比奥 feud.Barony = &博比奥BobbioBarony{}

func init() {
    f := BBobbio博比奥.(*博比奥BobbioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bobbio",
		TitleName: "博比奥",
		TitleCode: "b_bobbio",
	}
}
