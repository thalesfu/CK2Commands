package altay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 可可托海KoktokayBarony struct {
	feud.BaseBarony
}

var BKoktokay可可托海 feud.Barony = &可可托海KoktokayBarony{}

func init() {
    f := BKoktokay可可托海.(*可可托海KoktokayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "koktokay",
		TitleName: "可可托海",
		TitleCode: "b_koktokay",
	}
}
