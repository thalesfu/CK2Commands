package kandail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 耶摩那波梨YamanpalliBarony struct {
	feud.BaseBarony
}

var BYamanpalli耶摩那波梨 feud.Barony = &耶摩那波梨YamanpalliBarony{}

func init() {
	f := BYamanpalli耶摩那波梨.(*耶摩那波梨YamanpalliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yamanpalli",
		TitleName: "耶摩那波梨",
		TitleCode: "b_yamanpalli",
	}
}
