package somerset

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 凯里堡Castle_caryBarony struct {
	feud.BaseBarony
}

var BCastle_cary凯里堡 feud.Barony = &凯里堡Castle_caryBarony{}

func init() {
    f := BCastle_cary凯里堡.(*凯里堡Castle_caryBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castle_cary",
		TitleName: "凯里堡",
		TitleCode: "b_castle_cary",
	}
}
