package gallura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 多尔加利DorgaliBarony struct {
	feud.BaseBarony
}

var BDorgali多尔加利 feud.Barony = &多尔加利DorgaliBarony{}

func init() {
    f := BDorgali多尔加利.(*多尔加利DorgaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dorgali",
		TitleName: "多尔加利",
		TitleCode: "b_dorgali",
	}
}
