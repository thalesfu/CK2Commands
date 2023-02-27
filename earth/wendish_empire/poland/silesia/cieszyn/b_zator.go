package cieszyn

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎托尔ZatorBarony struct {
	feud.BaseBarony
}

var BZator扎托尔 feud.Barony = &扎托尔ZatorBarony{}

func init() {
    f := BZator扎托尔.(*扎托尔ZatorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zator",
		TitleName: "扎托尔",
		TitleCode: "b_zator",
	}
}
