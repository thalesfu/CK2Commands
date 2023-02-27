package lorraine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图勒ToulBarony struct {
	feud.BaseBarony
}

var BToul图勒 feud.Barony = &图勒ToulBarony{}

func init() {
    f := BToul图勒.(*图勒ToulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "toul",
		TitleName: "图勒",
		TitleCode: "b_toul",
	}
}
