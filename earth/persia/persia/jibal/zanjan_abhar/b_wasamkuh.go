package zanjan_abhar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦桑库赫WasamkuhBarony struct {
	feud.BaseBarony
}

var BWasamkuh瓦桑库赫 feud.Barony = &瓦桑库赫WasamkuhBarony{}

func init() {
    f := BWasamkuh瓦桑库赫.(*瓦桑库赫WasamkuhBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "wasamkuh",
		TitleName: "瓦桑库赫",
		TitleCode: "b_wasamkuh",
	}
}
