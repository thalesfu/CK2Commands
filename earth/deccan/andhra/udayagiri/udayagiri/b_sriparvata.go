package udayagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 室利钵伐多SriparvataBarony struct {
	feud.BaseBarony
}

var BSriparvata室利钵伐多 feud.Barony = &室利钵伐多SriparvataBarony{}

func init() {
    f := BSriparvata室利钵伐多.(*室利钵伐多SriparvataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sriparvata",
		TitleName: "室利钵伐多",
		TitleCode: "b_sriparvata",
	}
}
