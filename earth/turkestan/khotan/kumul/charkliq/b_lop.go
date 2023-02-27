package charkliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗卜LopBarony struct {
	feud.BaseBarony
}

var BLop罗卜 feud.Barony = &罗卜LopBarony{}

func init() {
    f := BLop罗卜.(*罗卜LopBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lop",
		TitleName: "罗卜",
		TitleCode: "b_lop",
	}
}
