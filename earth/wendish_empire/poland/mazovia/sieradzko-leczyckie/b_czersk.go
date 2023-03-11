package sieradzko-leczyckie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔斯克CzerskBarony struct {
	feud.BaseBarony
}

var BCzersk切尔斯克 feud.Barony = &切尔斯克CzerskBarony{}

func init() {
    f := BCzersk切尔斯克.(*切尔斯克CzerskBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "czersk",
		TitleName: "切尔斯克",
		TitleCode: "b_czersk",
	}
}
