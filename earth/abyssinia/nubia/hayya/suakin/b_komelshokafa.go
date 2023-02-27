package suakin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科梅尔索卡法KomelshokafaBarony struct {
	feud.BaseBarony
}

var BKomelshokafa科梅尔索卡法 feud.Barony = &科梅尔索卡法KomelshokafaBarony{}

func init() {
    f := BKomelshokafa科梅尔索卡法.(*科梅尔索卡法KomelshokafaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "komelshokafa",
		TitleName: "科梅尔索卡法",
		TitleCode: "b_komelshokafa",
	}
}
