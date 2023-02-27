package kantor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎托尔KantorBarony struct {
	feud.BaseBarony
}

var BKantor坎托尔 feud.Barony = &坎托尔KantorBarony{}

func init() {
    f := BKantor坎托尔.(*坎托尔KantorBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kantor",
		TitleName: "坎托尔",
		TitleCode: "b_kantor",
	}
}
