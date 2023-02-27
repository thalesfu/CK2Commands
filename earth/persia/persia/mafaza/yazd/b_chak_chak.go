package yazd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恰克恰克Chak_chakBarony struct {
	feud.BaseBarony
}

var BChak_chak恰克恰克 feud.Barony = &恰克恰克Chak_chakBarony{}

func init() {
    f := BChak_chak恰克恰克.(*恰克恰克Chak_chakBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chak_chak",
		TitleName: "恰克恰克",
		TitleCode: "b_chak_chak",
	}
}
