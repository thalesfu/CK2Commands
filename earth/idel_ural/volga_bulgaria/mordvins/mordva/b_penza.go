package mordva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奔萨PenzaBarony struct {
	feud.BaseBarony
}

var BPenza奔萨 feud.Barony = &奔萨PenzaBarony{}

func init() {
    f := BPenza奔萨.(*奔萨PenzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "penza",
		TitleName: "奔萨",
		TitleCode: "b_penza",
	}
}
