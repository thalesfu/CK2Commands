package asturias_de_santillana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 滨海桑蒂利亚纳SantillanadelmarBarony struct {
	feud.BaseBarony
}

var BSantillanadelmar滨海桑蒂利亚纳 feud.Barony = &滨海桑蒂利亚纳SantillanadelmarBarony{}

func init() {
    f := BSantillanadelmar滨海桑蒂利亚纳.(*滨海桑蒂利亚纳SantillanadelmarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "santillanadelmar",
		TitleName: "滨海桑蒂利亚纳",
		TitleCode: "b_santillanadelmar",
	}
}
