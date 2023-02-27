package saravan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨拉万SaravanBarony struct {
	feud.BaseBarony
}

var BSaravan萨拉万 feud.Barony = &萨拉万SaravanBarony{}

func init() {
    f := BSaravan萨拉万.(*萨拉万SaravanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "saravan",
		TitleName: "萨拉万",
		TitleCode: "b_saravan",
	}
}
