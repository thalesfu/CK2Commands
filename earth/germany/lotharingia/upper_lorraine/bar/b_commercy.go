package bar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科梅尔西CommercyBarony struct {
	feud.BaseBarony
}

var BCommercy科梅尔西 feud.Barony = &科梅尔西CommercyBarony{}

func init() {
    f := BCommercy科梅尔西.(*科梅尔西CommercyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "commercy",
		TitleName: "科梅尔西",
		TitleCode: "b_commercy",
	}
}
