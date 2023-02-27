package ouled_nail

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 布依拉TubirettBarony struct {
	feud.BaseBarony
}

var BTubirett布依拉 feud.Barony = &布依拉TubirettBarony{}

func init() {
    f := BTubirett布依拉.(*布依拉TubirettBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tubirett",
		TitleName: "布依拉",
		TitleCode: "b_tubirett",
	}
}
