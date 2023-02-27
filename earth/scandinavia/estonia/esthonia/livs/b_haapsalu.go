package livs

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈普萨卢HaapsaluBarony struct {
	feud.BaseBarony
}

var BHaapsalu哈普萨卢 feud.Barony = &哈普萨卢HaapsaluBarony{}

func init() {
    f := BHaapsalu哈普萨卢.(*哈普萨卢HaapsaluBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haapsalu",
		TitleName: "哈普萨卢",
		TitleCode: "b_haapsalu",
	}
}
