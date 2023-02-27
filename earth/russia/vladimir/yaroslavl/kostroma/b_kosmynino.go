package kostroma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯梅尼诺KosmyninoBarony struct {
	feud.BaseBarony
}

var BKosmynino科斯梅尼诺 feud.Barony = &科斯梅尼诺KosmyninoBarony{}

func init() {
    f := BKosmynino科斯梅尼诺.(*科斯梅尼诺KosmyninoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kosmynino",
		TitleName: "科斯梅尼诺",
		TitleCode: "b_kosmynino",
	}
}
