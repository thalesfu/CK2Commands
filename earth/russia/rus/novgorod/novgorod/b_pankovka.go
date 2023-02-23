package novgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 潘科夫卡PankovkaBarony struct {
	feud.BaseBarony
}

var BPankovka潘科夫卡 feud.Barony = &潘科夫卡PankovkaBarony{}

func init() {
	f := BPankovka潘科夫卡.(*潘科夫卡PankovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pankovka",
		TitleName: "潘科夫卡",
		TitleCode: "b_pankovka",
	}
}
