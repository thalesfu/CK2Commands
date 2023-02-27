package suenik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡拉恩GhapanBarony struct {
	feud.BaseBarony
}

var BGhapan卡拉恩 feud.Barony = &卡拉恩GhapanBarony{}

func init() {
    f := BGhapan卡拉恩.(*卡拉恩GhapanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ghapan",
		TitleName: "卡拉恩",
		TitleCode: "b_ghapan",
	}
}
