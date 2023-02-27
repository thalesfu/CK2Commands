package gerze

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吓嘎Jakar_bBarony struct {
	feud.BaseBarony
}

var BJakar_b吓嘎 feud.Barony = &吓嘎Jakar_bBarony{}

func init() {
    f := BJakar_b吓嘎.(*吓嘎Jakar_bBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "jakar_b",
		TitleName: "吓嘎",
		TitleCode: "b_jakar_b",
	}
}
