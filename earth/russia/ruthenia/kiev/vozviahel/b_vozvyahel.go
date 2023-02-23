package vozviahel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃兹维亚赫利VozvyahelBarony struct {
	feud.BaseBarony
}

var BVozvyahel沃兹维亚赫利 feud.Barony = &沃兹维亚赫利VozvyahelBarony{}

func init() {
	f := BVozvyahel沃兹维亚赫利.(*沃兹维亚赫利VozvyahelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vozvyahel",
		TitleName: "沃兹维亚赫利",
		TitleCode: "b_vozvyahel",
	}
}
