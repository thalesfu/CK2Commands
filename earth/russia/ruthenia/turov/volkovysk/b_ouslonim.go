package volkovysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌斯洛尼姆OuslonimBarony struct {
	feud.BaseBarony
}

var BOuslonim乌斯洛尼姆 feud.Barony = &乌斯洛尼姆OuslonimBarony{}

func init() {
    f := BOuslonim乌斯洛尼姆.(*乌斯洛尼姆OuslonimBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ouslonim",
		TitleName: "乌斯洛尼姆",
		TitleCode: "b_ouslonim",
	}
}
