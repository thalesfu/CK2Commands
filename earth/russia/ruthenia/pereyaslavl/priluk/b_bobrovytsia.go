package priluk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博布罗维齐亚BobrovytsiaBarony struct {
	feud.BaseBarony
}

var BBobrovytsia博布罗维齐亚 feud.Barony = &博布罗维齐亚BobrovytsiaBarony{}

func init() {
    f := BBobrovytsia博布罗维齐亚.(*博布罗维齐亚BobrovytsiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bobrovytsia",
		TitleName: "博布罗维齐亚",
		TitleCode: "b_bobrovytsia",
	}
}
