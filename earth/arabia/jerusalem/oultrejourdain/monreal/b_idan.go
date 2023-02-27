package monreal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊丹IdanBarony struct {
	feud.BaseBarony
}

var BIdan伊丹 feud.Barony = &伊丹IdanBarony{}

func init() {
    f := BIdan伊丹.(*伊丹IdanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "idan",
		TitleName: "伊丹",
		TitleCode: "b_idan",
	}
}
