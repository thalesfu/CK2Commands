package shrewsbury

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉德洛LudlowBarony struct {
	feud.BaseBarony
}

var BLudlow拉德洛 feud.Barony = &拉德洛LudlowBarony{}

func init() {
    f := BLudlow拉德洛.(*拉德洛LudlowBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ludlow",
		TitleName: "拉德洛",
		TitleCode: "b_ludlow",
	}
}
