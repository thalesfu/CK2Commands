package lomzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希温茨克SwieckBarony struct {
	feud.BaseBarony
}

var BSwieck希温茨克 feud.Barony = &希温茨克SwieckBarony{}

func init() {
    f := BSwieck希温茨克.(*希温茨克SwieckBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "swieck",
		TitleName: "希温茨克",
		TitleCode: "b_swieck",
	}
}
