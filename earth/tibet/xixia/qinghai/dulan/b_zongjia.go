package dulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 宗加ZongjiaBarony struct {
	feud.BaseBarony
}

var BZongjia宗加 feud.Barony = &宗加ZongjiaBarony{}

func init() {
	f := BZongjia宗加.(*宗加ZongjiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zongjia",
		TitleName: "宗加",
		TitleCode: "b_zongjia",
	}
}
