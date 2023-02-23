package traianopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂海罗TycheroBarony struct {
	feud.BaseBarony
}

var BTychero蒂海罗 feud.Barony = &蒂海罗TycheroBarony{}

func init() {
	f := BTychero蒂海罗.(*蒂海罗TycheroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tychero",
		TitleName: "蒂海罗",
		TitleCode: "b_tychero",
	}
}
