package begemder

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 杰克雷夫塔ChekereftaBarony struct {
	feud.BaseBarony
}

var BChekerefta杰克雷夫塔 feud.Barony = &杰克雷夫塔ChekereftaBarony{}

func init() {
    f := BChekerefta杰克雷夫塔.(*杰克雷夫塔ChekereftaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chekerefta",
		TitleName: "杰克雷夫塔",
		TitleCode: "b_chekerefta",
	}
}
