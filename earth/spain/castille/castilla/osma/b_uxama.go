package osma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌克萨马UxamaBarony struct {
	feud.BaseBarony
}

var BUxama乌克萨马 feud.Barony = &乌克萨马UxamaBarony{}

func init() {
	f := BUxama乌克萨马.(*乌克萨马UxamaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uxama",
		TitleName: "乌克萨马",
		TitleCode: "b_uxama",
	}
}
