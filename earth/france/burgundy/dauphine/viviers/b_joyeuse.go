package viviers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 茹瓦约斯JoyeuseBarony struct {
	feud.BaseBarony
}

var BJoyeuse茹瓦约斯 feud.Barony = &茹瓦约斯JoyeuseBarony{}

func init() {
    f := BJoyeuse茹瓦约斯.(*茹瓦约斯JoyeuseBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "joyeuse",
		TitleName: "茹瓦约斯",
		TitleCode: "b_joyeuse",
	}
}
