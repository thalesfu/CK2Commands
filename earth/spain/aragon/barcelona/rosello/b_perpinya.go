package rosello

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佩皮尼昂PerpinyaBarony struct {
	feud.BaseBarony
}

var BPerpinya佩皮尼昂 feud.Barony = &佩皮尼昂PerpinyaBarony{}

func init() {
    f := BPerpinya佩皮尼昂.(*佩皮尼昂PerpinyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "perpinya",
		TitleName: "佩皮尼昂",
		TitleCode: "b_perpinya",
	}
}
