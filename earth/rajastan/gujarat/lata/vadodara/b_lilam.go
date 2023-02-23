package vadodara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 梨蓝LilamBarony struct {
	feud.BaseBarony
}

var BLilam梨蓝 feud.Barony = &梨蓝LilamBarony{}

func init() {
	f := BLilam梨蓝.(*梨蓝LilamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lilam",
		TitleName: "梨蓝",
		TitleCode: "b_lilam",
	}
}
