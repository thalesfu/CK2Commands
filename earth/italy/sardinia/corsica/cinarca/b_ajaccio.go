package cinarca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿雅克肖AjaccioBarony struct {
	feud.BaseBarony
}

var BAjaccio阿雅克肖 feud.Barony = &阿雅克肖AjaccioBarony{}

func init() {
	f := BAjaccio阿雅克肖.(*阿雅克肖AjaccioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ajaccio",
		TitleName: "阿雅克肖",
		TitleCode: "b_ajaccio",
	}
}
