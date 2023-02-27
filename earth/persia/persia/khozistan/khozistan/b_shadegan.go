package khozistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙代甘ShadeganBarony struct {
	feud.BaseBarony
}

var BShadegan沙代甘 feud.Barony = &沙代甘ShadeganBarony{}

func init() {
    f := BShadegan沙代甘.(*沙代甘ShadeganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shadegan",
		TitleName: "沙代甘",
		TitleCode: "b_shadegan",
	}
}
