package medantaka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波招罗PachaurBarony struct {
	feud.BaseBarony
}

var BPachaur波招罗 feud.Barony = &波招罗PachaurBarony{}

func init() {
    f := BPachaur波招罗.(*波招罗PachaurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pachaur",
		TitleName: "波招罗",
		TitleCode: "b_pachaur",
	}
}
