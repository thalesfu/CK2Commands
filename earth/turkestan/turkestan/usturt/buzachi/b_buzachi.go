package buzachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布扎奇BuzachiBarony struct {
	feud.BaseBarony
}

var BBuzachi布扎奇 feud.Barony = &布扎奇BuzachiBarony{}

func init() {
    f := BBuzachi布扎奇.(*布扎奇BuzachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "buzachi",
		TitleName: "布扎奇",
		TitleCode: "b_buzachi",
	}
}
