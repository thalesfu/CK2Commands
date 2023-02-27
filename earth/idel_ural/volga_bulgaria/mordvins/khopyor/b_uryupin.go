package khopyor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌留平UryupinBarony struct {
	feud.BaseBarony
}

var BUryupin乌留平 feud.Barony = &乌留平UryupinBarony{}

func init() {
    f := BUryupin乌留平.(*乌留平UryupinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uryupin",
		TitleName: "乌留平",
		TitleCode: "b_uryupin",
	}
}
