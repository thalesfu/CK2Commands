package devon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托特尼斯TotnesBarony struct {
	feud.BaseBarony
}

var BTotnes托特尼斯 feud.Barony = &托特尼斯TotnesBarony{}

func init() {
	f := BTotnes托特尼斯.(*托特尼斯TotnesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "totnes",
		TitleName: "托特尼斯",
		TitleCode: "b_totnes",
	}
}
