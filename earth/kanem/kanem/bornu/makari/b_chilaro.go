package makari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇拉罗ChilaroBarony struct {
	feud.BaseBarony
}

var BChilaro奇拉罗 feud.Barony = &奇拉罗ChilaroBarony{}

func init() {
	f := BChilaro奇拉罗.(*奇拉罗ChilaroBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chilaro",
		TitleName: "奇拉罗",
		TitleCode: "b_chilaro",
	}
}
