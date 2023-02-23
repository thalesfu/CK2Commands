package nassau

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拿骚NassauBarony struct {
	feud.BaseBarony
}

var BNassau拿骚 feud.Barony = &拿骚NassauBarony{}

func init() {
	f := BNassau拿骚.(*拿骚NassauBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nassau",
		TitleName: "拿骚",
		TitleCode: "b_nassau",
	}
}
