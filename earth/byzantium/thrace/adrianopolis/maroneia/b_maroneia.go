package maroneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马戎尼亚MaroneiaBarony struct {
	feud.BaseBarony
}

var BMaroneia马戎尼亚 feud.Barony = &马戎尼亚MaroneiaBarony{}

func init() {
	f := BMaroneia马戎尼亚.(*马戎尼亚MaroneiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "maroneia",
		TitleName: "马戎尼亚",
		TitleCode: "b_maroneia",
	}
}
