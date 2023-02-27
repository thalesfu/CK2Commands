package sarkel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_顿涅茨基UstdonetskiyBarony struct {
	feud.BaseBarony
}

var BUstdonetskiy乌斯季_顿涅茨基 feud.Barony = &乌斯季_顿涅茨基UstdonetskiyBarony{}

func init() {
    f := BUstdonetskiy乌斯季_顿涅茨基.(*乌斯季_顿涅茨基UstdonetskiyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ustdonetskiy",
		TitleName: "乌斯季_顿涅茨基",
		TitleCode: "b_ustdonetskiy",
	}
}
