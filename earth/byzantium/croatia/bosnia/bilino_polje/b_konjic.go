package bilino_polje

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科尼茨KonjicBarony struct {
	feud.BaseBarony
}

var BKonjic科尼茨 feud.Barony = &科尼茨KonjicBarony{}

func init() {
    f := BKonjic科尼茨.(*科尼茨KonjicBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "konjic",
		TitleName: "科尼茨",
		TitleCode: "b_konjic",
	}
}
