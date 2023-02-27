package urbino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣马力诺SanmarinoBarony struct {
	feud.BaseBarony
}

var BSanmarino圣马力诺 feud.Barony = &圣马力诺SanmarinoBarony{}

func init() {
    f := BSanmarino圣马力诺.(*圣马力诺SanmarinoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sanmarino",
		TitleName: "圣马力诺",
		TitleCode: "b_sanmarino",
	}
}
