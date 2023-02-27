package masseniya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿特格贝AtegbeBarony struct {
	feud.BaseBarony
}

var BAtegbe阿特格贝 feud.Barony = &阿特格贝AtegbeBarony{}

func init() {
    f := BAtegbe阿特格贝.(*阿特格贝AtegbeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ategbe",
		TitleName: "阿特格贝",
		TitleCode: "b_ategbe",
	}
}
