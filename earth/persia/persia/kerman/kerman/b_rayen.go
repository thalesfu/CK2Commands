package kerman

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉延RayenBarony struct {
	feud.BaseBarony
}

var BRayen拉延 feud.Barony = &拉延RayenBarony{}

func init() {
    f := BRayen拉延.(*拉延RayenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rayen",
		TitleName: "拉延",
		TitleCode: "b_rayen",
	}
}
