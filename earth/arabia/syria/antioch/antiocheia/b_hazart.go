package antiocheia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈扎特HazartBarony struct {
	feud.BaseBarony
}

var BHazart哈扎特 feud.Barony = &哈扎特HazartBarony{}

func init() {
	f := BHazart哈扎特.(*哈扎特HazartBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hazart",
		TitleName: "哈扎特",
		TitleCode: "b_hazart",
	}
}
