package shirvan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨利杨SalyanBarony struct {
	feud.BaseBarony
}

var BSalyan萨利杨 feud.Barony = &萨利杨SalyanBarony{}

func init() {
    f := BSalyan萨利杨.(*萨利杨SalyanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "salyan",
		TitleName: "萨利杨",
		TitleCode: "b_salyan",
	}
}
