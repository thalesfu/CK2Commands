package assab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿萨布AssabBarony struct {
	feud.BaseBarony
}

var BAssab阿萨布 feud.Barony = &阿萨布AssabBarony{}

func init() {
    f := BAssab阿萨布.(*阿萨布AssabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "assab",
		TitleName: "阿萨布",
		TitleCode: "b_assab",
	}
}
