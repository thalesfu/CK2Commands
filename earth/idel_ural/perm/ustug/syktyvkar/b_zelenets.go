package syktyvkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泽列涅茨ZelenetsBarony struct {
	feud.BaseBarony
}

var BZelenets泽列涅茨 feud.Barony = &泽列涅茨ZelenetsBarony{}

func init() {
    f := BZelenets泽列涅茨.(*泽列涅茨ZelenetsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zelenets",
		TitleName: "泽列涅茨",
		TitleCode: "b_zelenets",
	}
}
