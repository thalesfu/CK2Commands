package warzazat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒂姆哈塔TimhataBarony struct {
	feud.BaseBarony
}

var BTimhata蒂姆哈塔 feud.Barony = &蒂姆哈塔TimhataBarony{}

func init() {
	f := BTimhata蒂姆哈塔.(*蒂姆哈塔TimhataBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "timhata",
		TitleName: "蒂姆哈塔",
		TitleCode: "b_timhata",
	}
}
