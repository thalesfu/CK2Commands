package sripuri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗什伐卢遮那Rajiva_lochanaBarony struct {
	feud.BaseBarony
}

var BRajiva_lochana罗什伐卢遮那 feud.Barony = &罗什伐卢遮那Rajiva_lochanaBarony{}

func init() {
    f := BRajiva_lochana罗什伐卢遮那.(*罗什伐卢遮那Rajiva_lochanaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "rajiva_lochana",
		TitleName: "罗什伐卢遮那",
		TitleCode: "b_rajiva_lochana",
	}
}
