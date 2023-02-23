package zoubtsov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 韦里吉诺VeriginoBarony struct {
	feud.BaseBarony
}

var BVerigino韦里吉诺 feud.Barony = &韦里吉诺VeriginoBarony{}

func init() {
	f := BVerigino韦里吉诺.(*韦里吉诺VeriginoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "verigino",
		TitleName: "韦里吉诺",
		TitleCode: "b_verigino",
	}
}
