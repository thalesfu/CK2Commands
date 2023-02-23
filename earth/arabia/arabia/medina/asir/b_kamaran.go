package asir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡马兰KamaranBarony struct {
	feud.BaseBarony
}

var BKamaran卡马兰 feud.Barony = &卡马兰KamaranBarony{}

func init() {
	f := BKamaran卡马兰.(*卡马兰KamaranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kamaran",
		TitleName: "卡马兰",
		TitleCode: "b_kamaran",
	}
}
