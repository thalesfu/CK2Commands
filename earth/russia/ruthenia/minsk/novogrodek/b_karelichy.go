package novogrodek

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 科列利奇KarelichyBarony struct {
	feud.BaseBarony
}

var BKarelichy科列利奇 feud.Barony = &科列利奇KarelichyBarony{}

func init() {
    f := BKarelichy科列利奇.(*科列利奇KarelichyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "karelichy",
		TitleName: "科列利奇",
		TitleCode: "b_karelichy",
	}
}
