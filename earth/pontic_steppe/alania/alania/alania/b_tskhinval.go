package alania

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茨欣瓦利TskhinvalBarony struct {
	feud.BaseBarony
}

var BTskhinval茨欣瓦利 feud.Barony = &茨欣瓦利TskhinvalBarony{}

func init() {
    f := BTskhinval茨欣瓦利.(*茨欣瓦利TskhinvalBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tskhinval",
		TitleName: "茨欣瓦利",
		TitleCode: "b_tskhinval",
	}
}
