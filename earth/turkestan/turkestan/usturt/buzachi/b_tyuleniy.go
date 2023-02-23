package buzachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 秋列尼TyuleniyBarony struct {
	feud.BaseBarony
}

var BTyuleniy秋列尼 feud.Barony = &秋列尼TyuleniyBarony{}

func init() {
	f := BTyuleniy秋列尼.(*秋列尼TyuleniyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyuleniy",
		TitleName: "秋列尼",
		TitleCode: "b_tyuleniy",
	}
}
