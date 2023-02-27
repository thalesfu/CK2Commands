package kujawy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比得哥什BygdoszczBarony struct {
	feud.BaseBarony
}

var BBygdoszcz比得哥什 feud.Barony = &比得哥什BygdoszczBarony{}

func init() {
    f := BBygdoszcz比得哥什.(*比得哥什BygdoszczBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bygdoszcz",
		TitleName: "比得哥什",
		TitleCode: "b_bygdoszcz",
	}
}
