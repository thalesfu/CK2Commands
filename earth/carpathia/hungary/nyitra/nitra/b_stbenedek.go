package nitra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣拜奈代克StbenedekBarony struct {
	feud.BaseBarony
}

var BStbenedek圣拜奈代克 feud.Barony = &圣拜奈代克StbenedekBarony{}

func init() {
	f := BStbenedek圣拜奈代克.(*圣拜奈代克StbenedekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stbenedek",
		TitleName: "圣拜奈代克",
		TitleCode: "b_stbenedek",
	}
}
