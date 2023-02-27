package ob

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切卡尼哈ChekanikhaBarony struct {
	feud.BaseBarony
}

var BChekanikha切卡尼哈 feud.Barony = &切卡尼哈ChekanikhaBarony{}

func init() {
    f := BChekanikha切卡尼哈.(*切卡尼哈ChekanikhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chekanikha",
		TitleName: "切卡尼哈",
		TitleCode: "b_chekanikha",
	}
}
