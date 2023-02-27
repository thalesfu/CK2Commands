package egiin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图古勒TugulBarony struct {
	feud.BaseBarony
}

var BTugul图古勒 feud.Barony = &图古勒TugulBarony{}

func init() {
    f := BTugul图古勒.(*图古勒TugulBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tugul",
		TitleName: "图古勒",
		TitleCode: "b_tugul",
	}
}
