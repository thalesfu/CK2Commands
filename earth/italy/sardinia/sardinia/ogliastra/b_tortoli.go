package ogliastra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔托利TortoliBarony struct {
	feud.BaseBarony
}

var BTortoli托尔托利 feud.Barony = &托尔托利TortoliBarony{}

func init() {
	f := BTortoli托尔托利.(*托尔托利TortoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tortoli",
		TitleName: "托尔托利",
		TitleCode: "b_tortoli",
	}
}
