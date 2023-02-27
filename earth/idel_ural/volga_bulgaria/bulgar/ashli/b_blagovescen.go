package ashli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布拉戈维申BlagovescenBarony struct {
	feud.BaseBarony
}

var BBlagovescen布拉戈维申 feud.Barony = &布拉戈维申BlagovescenBarony{}

func init() {
    f := BBlagovescen布拉戈维申.(*布拉戈维申BlagovescenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "blagovescen",
		TitleName: "布拉戈维申",
		TitleCode: "b_blagovescen",
	}
}
