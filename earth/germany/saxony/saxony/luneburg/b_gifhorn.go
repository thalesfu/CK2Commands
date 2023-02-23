package luneburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉夫霍恩GifhornBarony struct {
	feud.BaseBarony
}

var BGifhorn吉夫霍恩 feud.Barony = &吉夫霍恩GifhornBarony{}

func init() {
	f := BGifhorn吉夫霍恩.(*吉夫霍恩GifhornBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gifhorn",
		TitleName: "吉夫霍恩",
		TitleCode: "b_gifhorn",
	}
}
