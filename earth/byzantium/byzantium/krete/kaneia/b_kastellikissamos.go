package kaneia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯泰利_基萨莫斯KastellikissamosBarony struct {
	feud.BaseBarony
}

var BKastellikissamos卡斯泰利_基萨莫斯 feud.Barony = &卡斯泰利_基萨莫斯KastellikissamosBarony{}

func init() {
    f := BKastellikissamos卡斯泰利_基萨莫斯.(*卡斯泰利_基萨莫斯KastellikissamosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kastellikissamos",
		TitleName: "卡斯泰利_基萨莫斯",
		TitleCode: "b_kastellikissamos",
	}
}
