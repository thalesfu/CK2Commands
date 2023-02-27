package chuvash

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科兹洛夫卡KozlovkaBarony struct {
	feud.BaseBarony
}

var BKozlovka科兹洛夫卡 feud.Barony = &科兹洛夫卡KozlovkaBarony{}

func init() {
    f := BKozlovka科兹洛夫卡.(*科兹洛夫卡KozlovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kozlovka",
		TitleName: "科兹洛夫卡",
		TitleCode: "b_kozlovka",
	}
}
