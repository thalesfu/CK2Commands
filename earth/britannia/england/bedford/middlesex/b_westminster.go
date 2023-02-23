package middlesex

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 威斯敏斯特WestminsterBarony struct {
	feud.BaseBarony
}

var BWestminster威斯敏斯特 feud.Barony = &威斯敏斯特WestminsterBarony{}

func init() {
	f := BWestminster威斯敏斯特.(*威斯敏斯特WestminsterBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "westminster",
		TitleName: "威斯敏斯特",
		TitleCode: "b_westminster",
	}
}
