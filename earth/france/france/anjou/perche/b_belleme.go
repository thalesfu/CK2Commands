package perche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贝莱姆BellemeBarony struct {
	feud.BaseBarony
}

var BBelleme贝莱姆 feud.Barony = &贝莱姆BellemeBarony{}

func init() {
	f := BBelleme贝莱姆.(*贝莱姆BellemeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belleme",
		TitleName: "贝莱姆",
		TitleCode: "b_belleme",
	}
}
