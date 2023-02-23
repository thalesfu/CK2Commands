package dunois

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克卢瓦CloyesBarony struct {
	feud.BaseBarony
}

var BCloyes克卢瓦 feud.Barony = &克卢瓦CloyesBarony{}

func init() {
	f := BCloyes克卢瓦.(*克卢瓦CloyesBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cloyes",
		TitleName: "克卢瓦",
		TitleCode: "b_cloyes",
	}
}
