package perugia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古比奥GubbioBarony struct {
	feud.BaseBarony
}

var BGubbio古比奥 feud.Barony = &古比奥GubbioBarony{}

func init() {
	f := BGubbio古比奥.(*古比奥GubbioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gubbio",
		TitleName: "古比奥",
		TitleCode: "b_gubbio",
	}
}
