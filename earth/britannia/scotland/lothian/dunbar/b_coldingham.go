package dunbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尔丁厄姆ColdinghamBarony struct {
	feud.BaseBarony
}

var BColdingham科尔丁厄姆 feud.Barony = &科尔丁厄姆ColdinghamBarony{}

func init() {
    f := BColdingham科尔丁厄姆.(*科尔丁厄姆ColdinghamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "coldingham",
		TitleName: "科尔丁厄姆",
		TitleCode: "b_coldingham",
	}
}
