package kalaus

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗兹利夫RozlivBarony struct {
	feud.BaseBarony
}

var BRozliv罗兹利夫 feud.Barony = &罗兹利夫RozlivBarony{}

func init() {
    f := BRozliv罗兹利夫.(*罗兹利夫RozlivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rozliv",
		TitleName: "罗兹利夫",
		TitleCode: "b_rozliv",
	}
}
