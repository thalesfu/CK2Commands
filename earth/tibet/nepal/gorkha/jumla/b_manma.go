package jumla

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 曼尼ManmaBarony struct {
	feud.BaseBarony
}

var BManma曼尼 feud.Barony = &曼尼ManmaBarony{}

func init() {
    f := BManma曼尼.(*曼尼ManmaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "manma",
		TitleName: "曼尼",
		TitleCode: "b_manma",
	}
}
