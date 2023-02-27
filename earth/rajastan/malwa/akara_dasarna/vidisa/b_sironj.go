package vidisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 希罗纳SironjBarony struct {
	feud.BaseBarony
}

var BSironj希罗纳 feud.Barony = &希罗纳SironjBarony{}

func init() {
    f := BSironj希罗纳.(*希罗纳SironjBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sironj",
		TitleName: "希罗纳",
		TitleCode: "b_sironj",
	}
}
