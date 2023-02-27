package chester

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 麦克尔斯菲尔德MacclesfieldBarony struct {
	feud.BaseBarony
}

var BMacclesfield麦克尔斯菲尔德 feud.Barony = &麦克尔斯菲尔德MacclesfieldBarony{}

func init() {
    f := BMacclesfield麦克尔斯菲尔德.(*麦克尔斯菲尔德MacclesfieldBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "macclesfield",
		TitleName: "麦克尔斯菲尔德",
		TitleCode: "b_macclesfield",
	}
}
