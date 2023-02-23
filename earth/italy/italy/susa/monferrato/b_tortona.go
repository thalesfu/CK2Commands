package monferrato

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔托纳TortonaBarony struct {
	feud.BaseBarony
}

var BTortona托尔托纳 feud.Barony = &托尔托纳TortonaBarony{}

func init() {
	f := BTortona托尔托纳.(*托尔托纳TortonaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tortona",
		TitleName: "托尔托纳",
		TitleCode: "b_tortona",
	}
}
