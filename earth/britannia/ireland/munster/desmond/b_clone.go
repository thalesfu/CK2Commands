package desmond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克朗CloneBarony struct {
	feud.BaseBarony
}

var BClone克朗 feud.Barony = &克朗CloneBarony{}

func init() {
    f := BClone克朗.(*克朗CloneBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "clone",
		TitleName: "克朗",
		TitleCode: "b_clone",
	}
}
