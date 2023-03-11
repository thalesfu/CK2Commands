package naro_fominsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳罗福明斯克Naro_fominskBarony struct {
	feud.BaseBarony
}

var BNaro_fominsk纳罗福明斯克 feud.Barony = &纳罗福明斯克Naro_fominskBarony{}

func init() {
    f := BNaro_fominsk纳罗福明斯克.(*纳罗福明斯克Naro_fominskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naro_fominsk",
		TitleName: "纳罗福明斯克",
		TitleCode: "b_naro-fominsk",
	}
}
