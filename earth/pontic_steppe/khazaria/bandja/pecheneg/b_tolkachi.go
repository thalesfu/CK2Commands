package pecheneg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔卡奇TolkachiBarony struct {
	feud.BaseBarony
}

var BTolkachi托尔卡奇 feud.Barony = &托尔卡奇TolkachiBarony{}

func init() {
    f := BTolkachi托尔卡奇.(*托尔卡奇TolkachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tolkachi",
		TitleName: "托尔卡奇",
		TitleCode: "b_tolkachi",
	}
}
