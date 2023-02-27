package mozhaysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌连UrenBarony struct {
	feud.BaseBarony
}

var BUren乌连 feud.Barony = &乌连UrenBarony{}

func init() {
    f := BUren乌连.(*乌连UrenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uren",
		TitleName: "乌连",
		TitleCode: "b_uren",
	}
}
