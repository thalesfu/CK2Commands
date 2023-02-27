package telemark

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 埃兹堡EidsborgBarony struct {
	feud.BaseBarony
}

var BEidsborg埃兹堡 feud.Barony = &埃兹堡EidsborgBarony{}

func init() {
    f := BEidsborg埃兹堡.(*埃兹堡EidsborgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "eidsborg",
		TitleName: "埃兹堡",
		TitleCode: "b_eidsborg",
	}
}
