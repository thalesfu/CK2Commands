package kolhapur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 鸠伦陀迦KurundakaBarony struct {
	feud.BaseBarony
}

var BKurundaka鸠伦陀迦 feud.Barony = &鸠伦陀迦KurundakaBarony{}

func init() {
	f := BKurundaka鸠伦陀迦.(*鸠伦陀迦KurundakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kurundaka",
		TitleName: "鸠伦陀迦",
		TitleCode: "b_kurundaka",
	}
}
