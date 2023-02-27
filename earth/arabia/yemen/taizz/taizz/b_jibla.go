package taizz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉卜拉JiblaBarony struct {
	feud.BaseBarony
}

var BJibla吉卜拉 feud.Barony = &吉卜拉JiblaBarony{}

func init() {
    f := BJibla吉卜拉.(*吉卜拉JiblaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jibla",
		TitleName: "吉卜拉",
		TitleCode: "b_jibla",
	}
}
