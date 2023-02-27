package cornouaille

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙托兰ChateaulinBarony struct {
	feud.BaseBarony
}

var BChateaulin沙托兰 feud.Barony = &沙托兰ChateaulinBarony{}

func init() {
    f := BChateaulin沙托兰.(*沙托兰ChateaulinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chateaulin",
		TitleName: "沙托兰",
		TitleCode: "b_chateaulin",
	}
}
