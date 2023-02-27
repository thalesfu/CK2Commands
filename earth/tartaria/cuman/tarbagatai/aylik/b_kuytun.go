package aylik

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奎屯KuytunBarony struct {
	feud.BaseBarony
}

var BKuytun奎屯 feud.Barony = &奎屯KuytunBarony{}

func init() {
    f := BKuytun奎屯.(*奎屯KuytunBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuytun",
		TitleName: "奎屯",
		TitleCode: "b_kuytun",
	}
}
