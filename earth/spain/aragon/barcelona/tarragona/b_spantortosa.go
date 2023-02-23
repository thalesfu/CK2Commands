package tarragona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 托尔托萨SpantortosaBarony struct {
	feud.BaseBarony
}

var BSpantortosa托尔托萨 feud.Barony = &托尔托萨SpantortosaBarony{}

func init() {
	f := BSpantortosa托尔托萨.(*托尔托萨SpantortosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "spantortosa",
		TitleName: "托尔托萨",
		TitleCode: "b_spantortosa",
	}
}
