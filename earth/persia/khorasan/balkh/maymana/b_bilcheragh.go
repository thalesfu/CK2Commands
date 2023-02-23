package maymana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比尔切拉格BilcheraghBarony struct {
	feud.BaseBarony
}

var BBilcheragh比尔切拉格 feud.Barony = &比尔切拉格BilcheraghBarony{}

func init() {
	f := BBilcheragh比尔切拉格.(*比尔切拉格BilcheraghBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilcheragh",
		TitleName: "比尔切拉格",
		TitleCode: "b_bilcheragh",
	}
}
