package khopyor

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯季_梅德韦季茨卡亚UstmedveditskayaBarony struct {
	feud.BaseBarony
}

var BUstmedveditskaya乌斯季_梅德韦季茨卡亚 feud.Barony = &乌斯季_梅德韦季茨卡亚UstmedveditskayaBarony{}

func init() {
    f := BUstmedveditskaya乌斯季_梅德韦季茨卡亚.(*乌斯季_梅德韦季茨卡亚UstmedveditskayaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ustmedveditskaya",
		TitleName: "乌斯季_梅德韦季茨卡亚",
		TitleCode: "b_ustmedveditskaya",
	}
}
