package kondana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗阇摩支RajmachiBarony struct {
	feud.BaseBarony
}

var BRajmachi罗阇摩支 feud.Barony = &罗阇摩支RajmachiBarony{}

func init() {
    f := BRajmachi罗阇摩支.(*罗阇摩支RajmachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajmachi",
		TitleName: "罗阇摩支",
		TitleCode: "b_rajmachi",
	}
}
