package bremen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 施塔德StadeBarony struct {
	feud.BaseBarony
}

var BStade施塔德 feud.Barony = &施塔德StadeBarony{}

func init() {
    f := BStade施塔德.(*施塔德StadeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "stade",
		TitleName: "施塔德",
		TitleCode: "b_stade",
	}
}
