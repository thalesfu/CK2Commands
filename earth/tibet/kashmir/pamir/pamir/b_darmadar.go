package pamir

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达麦大DarmadarBarony struct {
	feud.BaseBarony
}

var BDarmadar达麦大 feud.Barony = &达麦大DarmadarBarony{}

func init() {
    f := BDarmadar达麦大.(*达麦大DarmadarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "darmadar",
		TitleName: "达麦大",
		TitleCode: "b_darmadar",
	}
}
