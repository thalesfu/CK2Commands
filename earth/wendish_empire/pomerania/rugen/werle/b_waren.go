package werle

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦伦WarenBarony struct {
	feud.BaseBarony
}

var BWaren瓦伦 feud.Barony = &瓦伦WarenBarony{}

func init() {
    f := BWaren瓦伦.(*瓦伦WarenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "waren",
		TitleName: "瓦伦",
		TitleCode: "b_waren",
	}
}
