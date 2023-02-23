package achaia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈兰兹里察ChalandritzaBarony struct {
	feud.BaseBarony
}

var BChalandritza哈兰兹里察 feud.Barony = &哈兰兹里察ChalandritzaBarony{}

func init() {
	f := BChalandritza哈兰兹里察.(*哈兰兹里察ChalandritzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chalandritza",
		TitleName: "哈兰兹里察",
		TitleCode: "b_chalandritza",
	}
}
