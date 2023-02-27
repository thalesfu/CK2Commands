package tortosa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔托萨TortosaBarony struct {
	feud.BaseBarony
}

var BTortosa托尔托萨 feud.Barony = &托尔托萨TortosaBarony{}

func init() {
    f := BTortosa托尔托萨.(*托尔托萨TortosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tortosa",
		TitleName: "托尔托萨",
		TitleCode: "b_tortosa",
	}
}
