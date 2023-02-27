package diafunu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 吉迪马卡GuidimakaBarony struct {
	feud.BaseBarony
}

var BGuidimaka吉迪马卡 feud.Barony = &吉迪马卡GuidimakaBarony{}

func init() {
    f := BGuidimaka吉迪马卡.(*吉迪马卡GuidimakaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guidimaka",
		TitleName: "吉迪马卡",
		TitleCode: "b_guidimaka",
	}
}
