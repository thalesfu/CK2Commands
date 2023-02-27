package gurgan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈尔甘GurganBarony struct {
	feud.BaseBarony
}

var BGurgan戈尔甘 feud.Barony = &戈尔甘GurganBarony{}

func init() {
    f := BGurgan戈尔甘.(*戈尔甘GurganBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gurgan",
		TitleName: "戈尔甘",
		TitleCode: "b_gurgan",
	}
}
