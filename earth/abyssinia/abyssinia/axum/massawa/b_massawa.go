package massawa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 马萨瓦MassawaBarony struct {
	feud.BaseBarony
}

var BMassawa马萨瓦 feud.Barony = &马萨瓦MassawaBarony{}

func init() {
    f := BMassawa马萨瓦.(*马萨瓦MassawaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "massawa",
		TitleName: "马萨瓦",
		TitleCode: "b_massawa",
	}
}
