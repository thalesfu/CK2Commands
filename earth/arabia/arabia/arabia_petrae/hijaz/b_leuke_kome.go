package hijaz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卢克科梅Leuke_komeBarony struct {
	feud.BaseBarony
}

var BLeuke_kome卢克科梅 feud.Barony = &卢克科梅Leuke_komeBarony{}

func init() {
    f := BLeuke_kome卢克科梅.(*卢克科梅Leuke_komeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "leuke_kome",
		TitleName: "卢克科梅",
		TitleCode: "b_leuke_kome",
	}
}
