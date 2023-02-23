package syrmia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布尔格奈BurgenaeBarony struct {
	feud.BaseBarony
}

var BBurgenae布尔格奈 feud.Barony = &布尔格奈BurgenaeBarony{}

func init() {
	f := BBurgenae布尔格奈.(*布尔格奈BurgenaeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "burgenae",
		TitleName: "布尔格奈",
		TitleCode: "b_burgenae",
	}
}
