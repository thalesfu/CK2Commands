package bezhetsky_verh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 诺温卡NovinkaBarony struct {
	feud.BaseBarony
}

var BNovinka诺温卡 feud.Barony = &诺温卡NovinkaBarony{}

func init() {
    f := BNovinka诺温卡.(*诺温卡NovinkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "novinka",
		TitleName: "诺温卡",
		TitleCode: "b_novinka",
	}
}
