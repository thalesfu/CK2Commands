package minsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 扎斯拉夫尔ZaslawyeBarony struct {
	feud.BaseBarony
}

var BZaslawye扎斯拉夫尔 feud.Barony = &扎斯拉夫尔ZaslawyeBarony{}

func init() {
    f := BZaslawye扎斯拉夫尔.(*扎斯拉夫尔ZaslawyeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zaslawye",
		TitleName: "扎斯拉夫尔",
		TitleCode: "b_zaslawye",
	}
}
