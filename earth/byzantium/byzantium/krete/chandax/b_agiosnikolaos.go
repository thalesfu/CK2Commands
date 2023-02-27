package chandax

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣尼古劳斯AgiosnikolaosBarony struct {
	feud.BaseBarony
}

var BAgiosnikolaos圣尼古劳斯 feud.Barony = &圣尼古劳斯AgiosnikolaosBarony{}

func init() {
    f := BAgiosnikolaos圣尼古劳斯.(*圣尼古劳斯AgiosnikolaosBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "agiosnikolaos",
		TitleName: "圣尼古劳斯",
		TitleCode: "b_agiosnikolaos",
	}
}
