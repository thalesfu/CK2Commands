package kostroma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科斯特罗马KostromaBarony struct {
	feud.BaseBarony
}

var BKostroma科斯特罗马 feud.Barony = &科斯特罗马KostromaBarony{}

func init() {
	f := BKostroma科斯特罗马.(*科斯特罗马KostromaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kostroma",
		TitleName: "科斯特罗马",
		TitleCode: "b_kostroma",
	}
}
