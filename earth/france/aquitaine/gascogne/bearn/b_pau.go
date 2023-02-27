package bearn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波城PauBarony struct {
	feud.BaseBarony
}

var BPau波城 feud.Barony = &波城PauBarony{}

func init() {
    f := BPau波城.(*波城PauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pau",
		TitleName: "波城",
		TitleCode: "b_pau",
	}
}
