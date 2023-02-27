package torres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 博萨BosaBarony struct {
	feud.BaseBarony
}

var BBosa博萨 feud.Barony = &博萨BosaBarony{}

func init() {
    f := BBosa博萨.(*博萨BosaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bosa",
		TitleName: "博萨",
		TitleCode: "b_bosa",
	}
}
