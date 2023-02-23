package jerusalem

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙吉萨MontgisardBarony struct {
	feud.BaseBarony
}

var BMontgisard蒙吉萨 feud.Barony = &蒙吉萨MontgisardBarony{}

func init() {
	f := BMontgisard蒙吉萨.(*蒙吉萨MontgisardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montgisard",
		TitleName: "蒙吉萨",
		TitleCode: "b_montgisard",
	}
}
