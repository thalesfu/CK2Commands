package traungau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 林茨LinzBarony struct {
	feud.BaseBarony
}

var BLinz林茨 feud.Barony = &林茨LinzBarony{}

func init() {
    f := BLinz林茨.(*林茨LinzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "linz",
		TitleName: "林茨",
		TitleCode: "b_linz",
	}
}
