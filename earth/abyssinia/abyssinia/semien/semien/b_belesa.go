package semien

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比勒萨BelesaBarony struct {
	feud.BaseBarony
}

var BBelesa比勒萨 feud.Barony = &比勒萨BelesaBarony{}

func init() {
	f := BBelesa比勒萨.(*比勒萨BelesaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belesa",
		TitleName: "比勒萨",
		TitleCode: "b_belesa",
	}
}
