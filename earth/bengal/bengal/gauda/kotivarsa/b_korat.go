package kotivarsa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拘剌吒KoratBarony struct {
	feud.BaseBarony
}

var BKorat拘剌吒 feud.Barony = &拘剌吒KoratBarony{}

func init() {
    f := BKorat拘剌吒.(*拘剌吒KoratBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "korat",
		TitleName: "拘剌吒",
		TitleCode: "b_korat",
	}
}
