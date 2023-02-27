package kiev

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日托米尔ZhytomyrBarony struct {
	feud.BaseBarony
}

var BZhytomyr日托米尔 feud.Barony = &日托米尔ZhytomyrBarony{}

func init() {
    f := BZhytomyr日托米尔.(*日托米尔ZhytomyrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhytomyr",
		TitleName: "日托米尔",
		TitleCode: "b_zhytomyr",
	}
}
