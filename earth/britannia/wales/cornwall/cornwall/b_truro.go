package cornwall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特鲁罗TruroBarony struct {
	feud.BaseBarony
}

var BTruro特鲁罗 feud.Barony = &特鲁罗TruroBarony{}

func init() {
	f := BTruro特鲁罗.(*特鲁罗TruroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "truro",
		TitleName: "特鲁罗",
		TitleCode: "b_truro",
	}
}
