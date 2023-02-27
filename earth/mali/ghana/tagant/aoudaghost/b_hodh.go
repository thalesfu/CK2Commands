package aoudaghost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 霍德HodhBarony struct {
	feud.BaseBarony
}

var BHodh霍德 feud.Barony = &霍德HodhBarony{}

func init() {
    f := BHodh霍德.(*霍德HodhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "hodh",
		TitleName: "霍德",
		TitleCode: "b_hodh",
	}
}
