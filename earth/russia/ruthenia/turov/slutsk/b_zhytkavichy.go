package slutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日特科维奇ZhytkavichyBarony struct {
	feud.BaseBarony
}

var BZhytkavichy日特科维奇 feud.Barony = &日特科维奇ZhytkavichyBarony{}

func init() {
    f := BZhytkavichy日特科维奇.(*日特科维奇ZhytkavichyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhytkavichy",
		TitleName: "日特科维奇",
		TitleCode: "b_zhytkavichy",
	}
}
