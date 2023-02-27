package ramagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吠罗瞿罗BelagulaBarony struct {
	feud.BaseBarony
}

var BBelagula吠罗瞿罗 feud.Barony = &吠罗瞿罗BelagulaBarony{}

func init() {
    f := BBelagula吠罗瞿罗.(*吠罗瞿罗BelagulaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "belagula",
		TitleName: "吠罗瞿罗",
		TitleCode: "b_belagula",
	}
}
