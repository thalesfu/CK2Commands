package gorlitz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 黑尔恩胡特HerrhutBarony struct {
	feud.BaseBarony
}

var BHerrhut黑尔恩胡特 feud.Barony = &黑尔恩胡特HerrhutBarony{}

func init() {
    f := BHerrhut黑尔恩胡特.(*黑尔恩胡特HerrhutBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "herrhut",
		TitleName: "黑尔恩胡特",
		TitleCode: "b_herrhut",
	}
}
