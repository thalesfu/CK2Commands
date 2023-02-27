package yazd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚兹德YazdBarony struct {
	feud.BaseBarony
}

var BYazd亚兹德 feud.Barony = &亚兹德YazdBarony{}

func init() {
    f := BYazd亚兹德.(*亚兹德YazdBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yazd",
		TitleName: "亚兹德",
		TitleCode: "b_yazd",
	}
}
