package grassland_cheremisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奇斯塔伊CistayBarony struct {
	feud.BaseBarony
}

var BCistay奇斯塔伊 feud.Barony = &奇斯塔伊CistayBarony{}

func init() {
    f := BCistay奇斯塔伊.(*奇斯塔伊CistayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cistay",
		TitleName: "奇斯塔伊",
		TitleCode: "b_cistay",
	}
}
