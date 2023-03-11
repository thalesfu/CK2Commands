package bryansk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 别日奇BezichiBarony struct {
	feud.BaseBarony
}

var BBezichi别日奇 feud.Barony = &别日奇BezichiBarony{}

func init() {
    f := BBezichi别日奇.(*别日奇BezichiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bezichi",
		TitleName: "别日奇",
		TitleCode: "b_bezichi",
	}
}
