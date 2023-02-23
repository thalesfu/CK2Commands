package nikaea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基奥斯KiosBarony struct {
	feud.BaseBarony
}

var BKios基奥斯 feud.Barony = &基奥斯KiosBarony{}

func init() {
	f := BKios基奥斯.(*基奥斯KiosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kios",
		TitleName: "基奥斯",
		TitleCode: "b_kios",
	}
}
