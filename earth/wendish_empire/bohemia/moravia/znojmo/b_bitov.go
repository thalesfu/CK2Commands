package znojmo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比托夫BitovBarony struct {
	feud.BaseBarony
}

var BBitov比托夫 feud.Barony = &比托夫BitovBarony{}

func init() {
    f := BBitov比托夫.(*比托夫BitovBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bitov",
		TitleName: "比托夫",
		TitleCode: "b_bitov",
	}
}
