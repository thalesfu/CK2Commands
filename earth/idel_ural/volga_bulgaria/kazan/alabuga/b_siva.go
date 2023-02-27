package alabuga

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡瓦SivaBarony struct {
	feud.BaseBarony
}

var BSiva锡瓦 feud.Barony = &锡瓦SivaBarony{}

func init() {
    f := BSiva锡瓦.(*锡瓦SivaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "siva",
		TitleName: "锡瓦",
		TitleCode: "b_siva",
	}
}
