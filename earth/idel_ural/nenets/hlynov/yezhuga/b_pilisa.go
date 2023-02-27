package yezhuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 皮利萨PilisaBarony struct {
	feud.BaseBarony
}

var BPilisa皮利萨 feud.Barony = &皮利萨PilisaBarony{}

func init() {
    f := BPilisa皮利萨.(*皮利萨PilisaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pilisa",
		TitleName: "皮利萨",
		TitleCode: "b_pilisa",
	}
}
