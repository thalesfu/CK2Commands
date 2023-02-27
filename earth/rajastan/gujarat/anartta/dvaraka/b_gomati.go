package dvaraka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈摩蒂GomatiBarony struct {
	feud.BaseBarony
}

var BGomati戈摩蒂 feud.Barony = &戈摩蒂GomatiBarony{}

func init() {
    f := BGomati戈摩蒂.(*戈摩蒂GomatiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gomati",
		TitleName: "戈摩蒂",
		TitleCode: "b_gomati",
	}
}
