package onega

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 普里亚扎PryazhaBarony struct {
	feud.BaseBarony
}

var BPryazha普里亚扎 feud.Barony = &普里亚扎PryazhaBarony{}

func init() {
	f := BPryazha普里亚扎.(*普里亚扎PryazhaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pryazha",
		TitleName: "普里亚扎",
		TitleCode: "b_pryazha",
	}
}
