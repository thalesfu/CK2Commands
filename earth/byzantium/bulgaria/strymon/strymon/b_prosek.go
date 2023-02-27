package strymon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普罗塞克ProsekBarony struct {
	feud.BaseBarony
}

var BProsek普罗塞克 feud.Barony = &普罗塞克ProsekBarony{}

func init() {
    f := BProsek普罗塞克.(*普罗塞克ProsekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "prosek",
		TitleName: "普罗塞克",
		TitleCode: "b_prosek",
	}
}
