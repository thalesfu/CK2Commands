package ile_de_france

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 巴黎ParisBarony struct {
	feud.BaseBarony
}

var BParis巴黎 feud.Barony = &巴黎ParisBarony{}

func init() {
    f := BParis巴黎.(*巴黎ParisBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "paris",
		TitleName: "巴黎",
		TitleCode: "b_paris",
	}
}
