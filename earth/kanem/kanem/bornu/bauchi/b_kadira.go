package bauchi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡迪拉KadiraBarony struct {
	feud.BaseBarony
}

var BKadira卡迪拉 feud.Barony = &卡迪拉KadiraBarony{}

func init() {
	f := BKadira卡迪拉.(*卡迪拉KadiraBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kadira",
		TitleName: "卡迪拉",
		TitleCode: "b_kadira",
	}
}
