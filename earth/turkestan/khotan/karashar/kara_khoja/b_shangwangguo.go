package kara_khoja

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 山王国ShangwangguoBarony struct {
	feud.BaseBarony
}

var BShangwangguo山王国 feud.Barony = &山王国ShangwangguoBarony{}

func init() {
    f := BShangwangguo山王国.(*山王国ShangwangguoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "shangwangguo",
		TitleName: "山王国",
		TitleCode: "b_shangwangguo",
	}
}
