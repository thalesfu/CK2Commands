package najera

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 纳赫拉NajeraBarony struct {
	feud.BaseBarony
}

var BNajera纳赫拉 feud.Barony = &纳赫拉NajeraBarony{}

func init() {
    f := BNajera纳赫拉.(*纳赫拉NajeraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "najera",
		TitleName: "纳赫拉",
		TitleCode: "b_najera",
	}
}
