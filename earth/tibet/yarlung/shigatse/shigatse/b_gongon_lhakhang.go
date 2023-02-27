package shigatse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贵恩拉康Gongon_lhakhangBarony struct {
	feud.BaseBarony
}

var BGongon_lhakhang贵恩拉康 feud.Barony = &贵恩拉康Gongon_lhakhangBarony{}

func init() {
    f := BGongon_lhakhang贵恩拉康.(*贵恩拉康Gongon_lhakhangBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gongon_lhakhang",
		TitleName: "贵恩拉康",
		TitleCode: "b_gongon_lhakhang",
	}
}
