package terkhiin_tsagaan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔里亚特TariatBarony struct {
	feud.BaseBarony
}

var BTariat塔里亚特 feud.Barony = &塔里亚特TariatBarony{}

func init() {
    f := BTariat塔里亚特.(*塔里亚特TariatBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tariat",
		TitleName: "塔里亚特",
		TitleCode: "b_tariat",
	}
}
