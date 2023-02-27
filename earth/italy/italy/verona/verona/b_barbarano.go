package verona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴尔巴拉诺BarbaranoBarony struct {
	feud.BaseBarony
}

var BBarbarano巴尔巴拉诺 feud.Barony = &巴尔巴拉诺BarbaranoBarony{}

func init() {
    f := BBarbarano巴尔巴拉诺.(*巴尔巴拉诺BarbaranoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "barbarano",
		TitleName: "巴尔巴拉诺",
		TitleCode: "b_barbarano",
	}
}
