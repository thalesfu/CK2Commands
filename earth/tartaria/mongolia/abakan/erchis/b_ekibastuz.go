package erchis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃基巴斯图兹EkibastuzBarony struct {
	feud.BaseBarony
}

var BEkibastuz埃基巴斯图兹 feud.Barony = &埃基巴斯图兹EkibastuzBarony{}

func init() {
    f := BEkibastuz埃基巴斯图兹.(*埃基巴斯图兹EkibastuzBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ekibastuz",
		TitleName: "埃基巴斯图兹",
		TitleCode: "b_ekibastuz",
	}
}
