package niederbayern

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 圣恩尔马尔Sankt_englmarBarony struct {
	feud.BaseBarony
}

var BSankt_englmar圣恩尔马尔 feud.Barony = &圣恩尔马尔Sankt_englmarBarony{}

func init() {
    f := BSankt_englmar圣恩尔马尔.(*圣恩尔马尔Sankt_englmarBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sankt_englmar",
		TitleName: "圣恩尔马尔",
		TitleCode: "b_sankt_englmar",
	}
}
