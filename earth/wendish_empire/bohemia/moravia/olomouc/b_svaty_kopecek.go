package olomouc

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣科佩切克Svaty_kopecekBarony struct {
	feud.BaseBarony
}

var BSvaty_kopecek圣科佩切克 feud.Barony = &圣科佩切克Svaty_kopecekBarony{}

func init() {
    f := BSvaty_kopecek圣科佩切克.(*圣科佩切克Svaty_kopecekBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "svaty_kopecek",
		TitleName: "圣科佩切克",
		TitleCode: "b_svaty_kopecek",
	}
}
