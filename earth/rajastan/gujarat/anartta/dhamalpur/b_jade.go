package dhamalpur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阇提JadeBarony struct {
	feud.BaseBarony
}

var BJade阇提 feud.Barony = &阇提JadeBarony{}

func init() {
    f := BJade阇提.(*阇提JadeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jade",
		TitleName: "阇提",
		TitleCode: "b_jade",
	}
}
