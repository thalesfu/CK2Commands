package torzhok

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔若克TorzhokBarony struct {
	feud.BaseBarony
}

var BTorzhok托尔若克 feud.Barony = &托尔若克TorzhokBarony{}

func init() {
    f := BTorzhok托尔若克.(*托尔若克TorzhokBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "torzhok",
		TitleName: "托尔若克",
		TitleCode: "b_torzhok",
	}
}
