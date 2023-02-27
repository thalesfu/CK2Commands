package brandenburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 于特博格JuterborgBarony struct {
	feud.BaseBarony
}

var BJuterborg于特博格 feud.Barony = &于特博格JuterborgBarony{}

func init() {
    f := BJuterborg于特博格.(*于特博格JuterborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "juterborg",
		TitleName: "于特博格",
		TitleCode: "b_juterborg",
	}
}
