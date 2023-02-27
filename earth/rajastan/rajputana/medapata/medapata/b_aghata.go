package medapata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿伽吒AghataBarony struct {
	feud.BaseBarony
}

var BAghata阿伽吒 feud.Barony = &阿伽吒AghataBarony{}

func init() {
    f := BAghata阿伽吒.(*阿伽吒AghataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aghata",
		TitleName: "阿伽吒",
		TitleCode: "b_aghata",
	}
}
