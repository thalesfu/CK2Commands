package keriya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 周村ZhoucunBarony struct {
	feud.BaseBarony
}

var BZhoucun周村 feud.Barony = &周村ZhoucunBarony{}

func init() {
    f := BZhoucun周村.(*周村ZhoucunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zhoucun",
		TitleName: "周村",
		TitleCode: "b_zhoucun",
	}
}
