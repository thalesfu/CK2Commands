package blois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢瓦尔河畔肖蒙ChaumontsurloireBarony struct {
	feud.BaseBarony
}

var BChaumontsurloire卢瓦尔河畔肖蒙 feud.Barony = &卢瓦尔河畔肖蒙ChaumontsurloireBarony{}

func init() {
	f := BChaumontsurloire卢瓦尔河畔肖蒙.(*卢瓦尔河畔肖蒙ChaumontsurloireBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chaumontsurloire",
		TitleName: "卢瓦尔河畔肖蒙",
		TitleCode: "b_chaumontsurloire",
	}
}
