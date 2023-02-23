package trier

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比特堡BitburgBarony struct {
	feud.BaseBarony
}

var BBitburg比特堡 feud.Barony = &比特堡BitburgBarony{}

func init() {
	f := BBitburg比特堡.(*比特堡BitburgBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bitburg",
		TitleName: "比特堡",
		TitleCode: "b_bitburg",
	}
}
