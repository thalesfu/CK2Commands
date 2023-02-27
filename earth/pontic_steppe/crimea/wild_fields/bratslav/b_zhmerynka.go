package bratslav

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 日梅林卡ZhmerynkaBarony struct {
	feud.BaseBarony
}

var BZhmerynka日梅林卡 feud.Barony = &日梅林卡ZhmerynkaBarony{}

func init() {
    f := BZhmerynka日梅林卡.(*日梅林卡ZhmerynkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhmerynka",
		TitleName: "日梅林卡",
		TitleCode: "b_zhmerynka",
	}
}
