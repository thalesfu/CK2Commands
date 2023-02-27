package itil

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈拉巴利KharabaliBarony struct {
	feud.BaseBarony
}

var BKharabali哈拉巴利 feud.Barony = &哈拉巴利KharabaliBarony{}

func init() {
    f := BKharabali哈拉巴利.(*哈拉巴利KharabaliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kharabali",
		TitleName: "哈拉巴利",
		TitleCode: "b_kharabali",
	}
}
