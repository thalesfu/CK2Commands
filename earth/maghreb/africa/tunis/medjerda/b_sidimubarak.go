package medjerda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 西迪穆巴拉克SidimubarakBarony struct {
	feud.BaseBarony
}

var BSidimubarak西迪穆巴拉克 feud.Barony = &西迪穆巴拉克SidimubarakBarony{}

func init() {
    f := BSidimubarak西迪穆巴拉克.(*西迪穆巴拉克SidimubarakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sidimubarak",
		TitleName: "西迪穆巴拉克",
		TitleCode: "b_sidimubarak",
	}
}
