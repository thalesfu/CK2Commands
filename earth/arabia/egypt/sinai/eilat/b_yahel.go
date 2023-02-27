package eilat

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 雅赫拉YahelBarony struct {
	feud.BaseBarony
}

var BYahel雅赫拉 feud.Barony = &雅赫拉YahelBarony{}

func init() {
    f := BYahel雅赫拉.(*雅赫拉YahelBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yahel",
		TitleName: "雅赫拉",
		TitleCode: "b_yahel",
	}
}
