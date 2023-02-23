package alexandretta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨凡提卡SarvantikarBarony struct {
	feud.BaseBarony
}

var BSarvantikar萨凡提卡 feud.Barony = &萨凡提卡SarvantikarBarony{}

func init() {
	f := BSarvantikar萨凡提卡.(*萨凡提卡SarvantikarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarvantikar",
		TitleName: "萨凡提卡",
		TitleCode: "b_sarvantikar",
	}
}
