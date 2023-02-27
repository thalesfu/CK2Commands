package gegyai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雄巴ZhungpaBarony struct {
	feud.BaseBarony
}

var BZhungpa雄巴 feud.Barony = &雄巴ZhungpaBarony{}

func init() {
    f := BZhungpa雄巴.(*雄巴ZhungpaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhungpa",
		TitleName: "雄巴",
		TitleCode: "b_zhungpa",
	}
}
