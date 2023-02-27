package lomzynska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沃姆扎LomzaBarony struct {
	feud.BaseBarony
}

var BLomza沃姆扎 feud.Barony = &沃姆扎LomzaBarony{}

func init() {
    f := BLomza沃姆扎.(*沃姆扎LomzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lomza",
		TitleName: "沃姆扎",
		TitleCode: "b_lomza",
	}
}
