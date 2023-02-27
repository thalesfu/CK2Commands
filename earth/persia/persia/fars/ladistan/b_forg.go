package ladistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 佛拉格ForgBarony struct {
	feud.BaseBarony
}

var BForg佛拉格 feud.Barony = &佛拉格ForgBarony{}

func init() {
    f := BForg佛拉格.(*佛拉格ForgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "forg",
		TitleName: "佛拉格",
		TitleCode: "b_forg",
	}
}
