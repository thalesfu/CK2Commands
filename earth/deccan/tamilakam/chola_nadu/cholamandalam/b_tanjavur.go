package cholamandalam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 檀伽瓦TanjavurBarony struct {
	feud.BaseBarony
}

var BTanjavur檀伽瓦 feud.Barony = &檀伽瓦TanjavurBarony{}

func init() {
    f := BTanjavur檀伽瓦.(*檀伽瓦TanjavurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tanjavur",
		TitleName: "檀伽瓦",
		TitleCode: "b_tanjavur",
	}
}
