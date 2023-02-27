package bobo_dyulasso

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔戈拉KongolaBarony struct {
	feud.BaseBarony
}

var BKongola孔戈拉 feud.Barony = &孔戈拉KongolaBarony{}

func init() {
    f := BKongola孔戈拉.(*孔戈拉KongolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kongola",
		TitleName: "孔戈拉",
		TitleCode: "b_kongola",
	}
}
