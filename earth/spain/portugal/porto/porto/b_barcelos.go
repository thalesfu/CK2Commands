package porto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴塞卢什BarcelosBarony struct {
	feud.BaseBarony
}

var BBarcelos巴塞卢什 feud.Barony = &巴塞卢什BarcelosBarony{}

func init() {
	f := BBarcelos巴塞卢什.(*巴塞卢什BarcelosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barcelos",
		TitleName: "巴塞卢什",
		TitleCode: "b_barcelos",
	}
}
