package poitiers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 帕尔特奈ParthenayBarony struct {
	feud.BaseBarony
}

var BParthenay帕尔特奈 feud.Barony = &帕尔特奈ParthenayBarony{}

func init() {
    f := BParthenay帕尔特奈.(*帕尔特奈ParthenayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "parthenay",
		TitleName: "帕尔特奈",
		TitleCode: "b_parthenay",
	}
}
