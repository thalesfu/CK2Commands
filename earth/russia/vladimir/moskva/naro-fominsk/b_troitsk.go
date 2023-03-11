package naro-fominsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 特罗伊茨克TroitskBarony struct {
	feud.BaseBarony
}

var BTroitsk特罗伊茨克 feud.Barony = &特罗伊茨克TroitskBarony{}

func init() {
    f := BTroitsk特罗伊茨克.(*特罗伊茨克TroitskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "troitsk",
		TitleName: "特罗伊茨克",
		TitleCode: "b_troitsk",
	}
}
