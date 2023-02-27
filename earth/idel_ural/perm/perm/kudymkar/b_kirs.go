package kudymkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 基尔斯KirsBarony struct {
	feud.BaseBarony
}

var BKirs基尔斯 feud.Barony = &基尔斯KirsBarony{}

func init() {
    f := BKirs基尔斯.(*基尔斯KirsBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kirs",
		TitleName: "基尔斯",
		TitleCode: "b_kirs",
	}
}
