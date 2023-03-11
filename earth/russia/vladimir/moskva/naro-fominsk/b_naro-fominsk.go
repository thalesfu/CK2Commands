package naro-fominsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳罗福明斯克Naro-fominskBarony struct {
	feud.BaseBarony
}

var BNaro-fominsk纳罗福明斯克 feud.Barony = &纳罗福明斯克Naro-fominskBarony{}

func init() {
    f := BNaro-fominsk纳罗福明斯克.(*纳罗福明斯克Naro-fominskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naro-fominsk",
		TitleName: "纳罗福明斯克",
		TitleCode: "b_naro-fominsk",
	}
}
