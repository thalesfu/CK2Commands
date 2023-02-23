package wurttemberg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗伊特林根ReutlingenBarony struct {
	feud.BaseBarony
}

var BReutlingen罗伊特林根 feud.Barony = &罗伊特林根ReutlingenBarony{}

func init() {
	f := BReutlingen罗伊特林根.(*罗伊特林根ReutlingenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "reutlingen",
		TitleName: "罗伊特林根",
		TitleCode: "b_reutlingen",
	}
}
