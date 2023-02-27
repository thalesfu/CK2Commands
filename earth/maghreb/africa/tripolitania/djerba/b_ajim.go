package djerba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 艾吉姆AjimBarony struct {
	feud.BaseBarony
}

var BAjim艾吉姆 feud.Barony = &艾吉姆AjimBarony{}

func init() {
    f := BAjim艾吉姆.(*艾吉姆AjimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ajim",
		TitleName: "艾吉姆",
		TitleCode: "b_ajim",
	}
}
