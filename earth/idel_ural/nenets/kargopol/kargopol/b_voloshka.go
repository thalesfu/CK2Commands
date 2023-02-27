package kargopol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃洛什卡VoloshkaBarony struct {
	feud.BaseBarony
}

var BVoloshka沃洛什卡 feud.Barony = &沃洛什卡VoloshkaBarony{}

func init() {
    f := BVoloshka沃洛什卡.(*沃洛什卡VoloshkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "voloshka",
		TitleName: "沃洛什卡",
		TitleCode: "b_voloshka",
	}
}
