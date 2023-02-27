package oleshye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡缅斯克KamianskBarony struct {
	feud.BaseBarony
}

var BKamiansk卡缅斯克 feud.Barony = &卡缅斯克KamianskBarony{}

func init() {
    f := BKamiansk卡缅斯克.(*卡缅斯克KamianskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamiansk",
		TitleName: "卡缅斯克",
		TitleCode: "b_kamiansk",
	}
}
