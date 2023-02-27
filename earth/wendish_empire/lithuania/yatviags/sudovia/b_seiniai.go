package sudovia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塞尼艾SeiniaiBarony struct {
	feud.BaseBarony
}

var BSeiniai塞尼艾 feud.Barony = &塞尼艾SeiniaiBarony{}

func init() {
    f := BSeiniai塞尼艾.(*塞尼艾SeiniaiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seiniai",
		TitleName: "塞尼艾",
		TitleCode: "b_seiniai",
	}
}
