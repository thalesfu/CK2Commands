package smolensk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯摩棱斯克SmolenskBarony struct {
	feud.BaseBarony
}

var BSmolensk斯摩棱斯克 feud.Barony = &斯摩棱斯克SmolenskBarony{}

func init() {
	f := BSmolensk斯摩棱斯克.(*斯摩棱斯克SmolenskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "smolensk",
		TitleName: "斯摩棱斯克",
		TitleCode: "b_smolensk",
	}
}
