package bezichy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔马奇TolmachiBarony struct {
	feud.BaseBarony
}

var BTolmachi托尔马奇 feud.Barony = &托尔马奇TolmachiBarony{}

func init() {
    f := BTolmachi托尔马奇.(*托尔马奇TolmachiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tolmachi",
		TitleName: "托尔马奇",
		TitleCode: "b_tolmachi",
	}
}
