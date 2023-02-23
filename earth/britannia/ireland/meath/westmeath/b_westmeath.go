package westmeath

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦斯特米斯WestmeathBarony struct {
	feud.BaseBarony
}

var BWestmeath韦斯特米斯 feud.Barony = &韦斯特米斯WestmeathBarony{}

func init() {
	f := BWestmeath韦斯特米斯.(*韦斯特米斯WestmeathBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "westmeath",
		TitleName: "韦斯特米斯",
		TitleCode: "b_westmeath",
	}
}
