package arezzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣塞波尔克罗SansepolcroBarony struct {
	feud.BaseBarony
}

var BSansepolcro圣塞波尔克罗 feud.Barony = &圣塞波尔克罗SansepolcroBarony{}

func init() {
	f := BSansepolcro圣塞波尔克罗.(*圣塞波尔克罗SansepolcroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sansepolcro",
		TitleName: "圣塞波尔克罗",
		TitleCode: "b_sansepolcro",
	}
}
