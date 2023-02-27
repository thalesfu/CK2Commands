package limousin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙吕ChalusBarony struct {
	feud.BaseBarony
}

var BChalus沙吕 feud.Barony = &沙吕ChalusBarony{}

func init() {
    f := BChalus沙吕.(*沙吕ChalusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chalus",
		TitleName: "沙吕",
		TitleCode: "b_chalus",
	}
}
