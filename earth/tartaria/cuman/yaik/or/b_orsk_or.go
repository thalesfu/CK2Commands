package or

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥尔斯克Orsk_orBarony struct {
	feud.BaseBarony
}

var BOrsk_or奥尔斯克 feud.Barony = &奥尔斯克Orsk_orBarony{}

func init() {
    f := BOrsk_or奥尔斯克.(*奥尔斯克Orsk_orBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "orsk_or",
		TitleName: "奥尔斯克",
		TitleCode: "b_orsk_or",
	}
}
