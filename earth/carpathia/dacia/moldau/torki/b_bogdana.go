package torki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博格丹纳BogdanaBarony struct {
	feud.BaseBarony
}

var BBogdana博格丹纳 feud.Barony = &博格丹纳BogdanaBarony{}

func init() {
    f := BBogdana博格丹纳.(*博格丹纳BogdanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bogdana",
		TitleName: "博格丹纳",
		TitleCode: "b_bogdana",
	}
}
