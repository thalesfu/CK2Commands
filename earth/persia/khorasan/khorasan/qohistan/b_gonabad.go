package qohistan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 戈纳巴德GonabadBarony struct {
	feud.BaseBarony
}

var BGonabad戈纳巴德 feud.Barony = &戈纳巴德GonabadBarony{}

func init() {
    f := BGonabad戈纳巴德.(*戈纳巴德GonabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gonabad",
		TitleName: "戈纳巴德",
		TitleCode: "b_gonabad",
	}
}
