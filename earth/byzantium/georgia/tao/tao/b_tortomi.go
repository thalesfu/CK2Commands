package tao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔图姆TortomiBarony struct {
	feud.BaseBarony
}

var BTortomi托尔图姆 feud.Barony = &托尔图姆TortomiBarony{}

func init() {
	f := BTortomi托尔图姆.(*托尔图姆TortomiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tortomi",
		TitleName: "托尔图姆",
		TitleCode: "b_tortomi",
	}
}
