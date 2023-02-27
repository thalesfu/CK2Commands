package kemi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯米耶尔维KemijarviBarony struct {
	feud.BaseBarony
}

var BKemijarvi凯米耶尔维 feud.Barony = &凯米耶尔维KemijarviBarony{}

func init() {
    f := BKemijarvi凯米耶尔维.(*凯米耶尔维KemijarviBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kemijarvi",
		TitleName: "凯米耶尔维",
		TitleCode: "b_kemijarvi",
	}
}
