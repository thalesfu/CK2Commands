package verona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 维罗纳VeronaBarony struct {
	feud.BaseBarony
}

var BVerona维罗纳 feud.Barony = &维罗纳VeronaBarony{}

func init() {
	f := BVerona维罗纳.(*维罗纳VeronaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verona",
		TitleName: "维罗纳",
		TitleCode: "b_verona",
	}
}
