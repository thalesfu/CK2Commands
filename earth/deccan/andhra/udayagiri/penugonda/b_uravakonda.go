package penugonda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 优罗婆军荼UravakondaBarony struct {
	feud.BaseBarony
}

var BUravakonda优罗婆军荼 feud.Barony = &优罗婆军荼UravakondaBarony{}

func init() {
    f := BUravakonda优罗婆军荼.(*优罗婆军荼UravakondaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "uravakonda",
		TitleName: "优罗婆军荼",
		TitleCode: "b_uravakonda",
	}
}
