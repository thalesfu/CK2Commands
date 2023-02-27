package baygal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌柳恩UlyunBarony struct {
	feud.BaseBarony
}

var BUlyun乌柳恩 feud.Barony = &乌柳恩UlyunBarony{}

func init() {
    f := BUlyun乌柳恩.(*乌柳恩UlyunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulyun",
		TitleName: "乌柳恩",
		TitleCode: "b_ulyun",
	}
}
