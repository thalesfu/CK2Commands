package qom

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎卡兰JamkaranBarony struct {
	feud.BaseBarony
}

var BJamkaran扎卡兰 feud.Barony = &扎卡兰JamkaranBarony{}

func init() {
	f := BJamkaran扎卡兰.(*扎卡兰JamkaranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jamkaran",
		TitleName: "扎卡兰",
		TitleCode: "b_jamkaran",
	}
}
