package penugonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉亚德加姆RayadurgamBarony struct {
	feud.BaseBarony
}

var BRayadurgam拉亚德加姆 feud.Barony = &拉亚德加姆RayadurgamBarony{}

func init() {
    f := BRayadurgam拉亚德加姆.(*拉亚德加姆RayadurgamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rayadurgam",
		TitleName: "拉亚德加姆",
		TitleCode: "b_rayadurgam",
	}
}
