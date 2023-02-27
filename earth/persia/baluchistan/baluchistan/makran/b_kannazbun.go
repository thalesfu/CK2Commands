package makran

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 坎纳兹本KannazbunBarony struct {
	feud.BaseBarony
}

var BKannazbun坎纳兹本 feud.Barony = &坎纳兹本KannazbunBarony{}

func init() {
    f := BKannazbun坎纳兹本.(*坎纳兹本KannazbunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kannazbun",
		TitleName: "坎纳兹本",
		TitleCode: "b_kannazbun",
	}
}
