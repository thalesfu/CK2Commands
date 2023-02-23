package barcelona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比克VicBarony struct {
	feud.BaseBarony
}

var BVic比克 feud.Barony = &比克VicBarony{}

func init() {
	f := BVic比克.(*比克VicBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vic",
		TitleName: "比克",
		TitleCode: "b_vic",
	}
}
