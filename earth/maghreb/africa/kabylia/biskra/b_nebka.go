package biskra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奈卜卡NebkaBarony struct {
	feud.BaseBarony
}

var BNebka奈卜卡 feud.Barony = &奈卜卡NebkaBarony{}

func init() {
    f := BNebka奈卜卡.(*奈卜卡NebkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nebka",
		TitleName: "奈卜卡",
		TitleCode: "b_nebka",
	}
}
