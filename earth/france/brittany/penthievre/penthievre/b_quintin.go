package penthievre

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎坦QuintinBarony struct {
	feud.BaseBarony
}

var BQuintin坎坦 feud.Barony = &坎坦QuintinBarony{}

func init() {
	f := BQuintin坎坦.(*坎坦QuintinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "quintin",
		TitleName: "坎坦",
		TitleCode: "b_quintin",
	}
}
