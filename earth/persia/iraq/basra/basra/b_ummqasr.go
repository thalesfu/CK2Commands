package basra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆盖萨尔UmmqasrBarony struct {
	feud.BaseBarony
}

var BUmmqasr乌姆盖萨尔 feud.Barony = &乌姆盖萨尔UmmqasrBarony{}

func init() {
    f := BUmmqasr乌姆盖萨尔.(*乌姆盖萨尔UmmqasrBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ummqasr",
		TitleName: "乌姆盖萨尔",
		TitleCode: "b_ummqasr",
	}
}
