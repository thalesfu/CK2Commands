package sarkel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨克尔SarkelBarony struct {
	feud.BaseBarony
}

var BSarkel萨克尔 feud.Barony = &萨克尔SarkelBarony{}

func init() {
    f := BSarkel萨克尔.(*萨克尔SarkelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sarkel",
		TitleName: "萨克尔",
		TitleCode: "b_sarkel",
	}
}
