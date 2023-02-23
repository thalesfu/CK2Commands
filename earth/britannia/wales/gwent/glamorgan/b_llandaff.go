package glamorgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兰达夫LlandaffBarony struct {
	feud.BaseBarony
}

var BLlandaff兰达夫 feud.Barony = &兰达夫LlandaffBarony{}

func init() {
	f := BLlandaff兰达夫.(*兰达夫LlandaffBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "llandaff",
		TitleName: "兰达夫",
		TitleCode: "b_llandaff",
	}
}
