package wadai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 古尔莫加GourmogaBarony struct {
	feud.BaseBarony
}

var BGourmoga古尔莫加 feud.Barony = &古尔莫加GourmogaBarony{}

func init() {
    f := BGourmoga古尔莫加.(*古尔莫加GourmogaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gourmoga",
		TitleName: "古尔莫加",
		TitleCode: "b_gourmoga",
	}
}
