package ferrara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科马基奥CommacchioBarony struct {
	feud.BaseBarony
}

var BCommacchio科马基奥 feud.Barony = &科马基奥CommacchioBarony{}

func init() {
    f := BCommacchio科马基奥.(*科马基奥CommacchioBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "commacchio",
		TitleName: "科马基奥",
		TitleCode: "b_commacchio",
	}
}
