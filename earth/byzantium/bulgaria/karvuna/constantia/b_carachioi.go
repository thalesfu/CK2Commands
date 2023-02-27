package constantia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉基奥CarachioiBarony struct {
	feud.BaseBarony
}

var BCarachioi卡拉基奥 feud.Barony = &卡拉基奥CarachioiBarony{}

func init() {
    f := BCarachioi卡拉基奥.(*卡拉基奥CarachioiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carachioi",
		TitleName: "卡拉基奥",
		TitleCode: "b_carachioi",
	}
}
