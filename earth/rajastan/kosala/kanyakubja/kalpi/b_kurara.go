package kalpi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 俱罗罗KuraraBarony struct {
	feud.BaseBarony
}

var BKurara俱罗罗 feud.Barony = &俱罗罗KuraraBarony{}

func init() {
	f := BKurara俱罗罗.(*俱罗罗KuraraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kurara",
		TitleName: "俱罗罗",
		TitleCode: "b_kurara",
	}
}
