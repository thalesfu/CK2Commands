package khiva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿塔亚布AtajabBarony struct {
	feud.BaseBarony
}

var BAtajab阿塔亚布 feud.Barony = &阿塔亚布AtajabBarony{}

func init() {
    f := BAtajab阿塔亚布.(*阿塔亚布AtajabBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "atajab",
		TitleName: "阿塔亚布",
		TitleCode: "b_atajab",
	}
}
