package vetluga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌连OurenBarony struct {
	feud.BaseBarony
}

var BOuren乌连 feud.Barony = &乌连OurenBarony{}

func init() {
    f := BOuren乌连.(*乌连OurenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouren",
		TitleName: "乌连",
		TitleCode: "b_ouren",
	}
}
