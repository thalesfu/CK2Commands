package bannu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科哈特KohatBarony struct {
	feud.BaseBarony
}

var BKohat科哈特 feud.Barony = &科哈特KohatBarony{}

func init() {
    f := BKohat科哈特.(*科哈特KohatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kohat",
		TitleName: "科哈特",
		TitleCode: "b_kohat",
	}
}
