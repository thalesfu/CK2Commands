package karur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 支室帝延ChishtianBarony struct {
	feud.BaseBarony
}

var BChishtian支室帝延 feud.Barony = &支室帝延ChishtianBarony{}

func init() {
    f := BChishtian支室帝延.(*支室帝延ChishtianBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chishtian",
		TitleName: "支室帝延",
		TitleCode: "b_chishtian",
	}
}
