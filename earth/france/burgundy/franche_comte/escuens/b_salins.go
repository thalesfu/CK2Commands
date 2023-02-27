package escuens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨兰SalinsBarony struct {
	feud.BaseBarony
}

var BSalins萨兰 feud.Barony = &萨兰SalinsBarony{}

func init() {
    f := BSalins萨兰.(*萨兰SalinsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salins",
		TitleName: "萨兰",
		TitleCode: "b_salins",
	}
}
