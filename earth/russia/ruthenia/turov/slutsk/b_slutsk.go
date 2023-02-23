package slutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 斯卢茨克SlutskBarony struct {
	feud.BaseBarony
}

var BSlutsk斯卢茨克 feud.Barony = &斯卢茨克SlutskBarony{}

func init() {
	f := BSlutsk斯卢茨克.(*斯卢茨克SlutskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "slutsk",
		TitleName: "斯卢茨克",
		TitleCode: "b_slutsk",
	}
}
