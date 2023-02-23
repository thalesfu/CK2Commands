package vakhan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吐蕃之门DaritubatBarony struct {
	feud.BaseBarony
}

var BDaritubat吐蕃之门 feud.Barony = &吐蕃之门DaritubatBarony{}

func init() {
	f := BDaritubat吐蕃之门.(*吐蕃之门DaritubatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "daritubat",
		TitleName: "吐蕃之门",
		TitleCode: "b_daritubat",
	}
}
