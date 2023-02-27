package ghana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴盖纳BaghenaBarony struct {
	feud.BaseBarony
}

var BBaghena巴盖纳 feud.Barony = &巴盖纳BaghenaBarony{}

func init() {
    f := BBaghena巴盖纳.(*巴盖纳BaghenaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "baghena",
		TitleName: "巴盖纳",
		TitleCode: "b_baghena",
	}
}
