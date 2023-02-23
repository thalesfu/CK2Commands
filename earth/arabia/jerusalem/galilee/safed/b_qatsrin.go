package safed

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡茨林QatsrinBarony struct {
	feud.BaseBarony
}

var BQatsrin卡茨林 feud.Barony = &卡茨林QatsrinBarony{}

func init() {
	f := BQatsrin卡茨林.(*卡茨林QatsrinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qatsrin",
		TitleName: "卡茨林",
		TitleCode: "b_qatsrin",
	}
}
