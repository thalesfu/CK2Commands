package napoli

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿夫拉戈拉AfragolaBarony struct {
	feud.BaseBarony
}

var BAfragola阿夫拉戈拉 feud.Barony = &阿夫拉戈拉AfragolaBarony{}

func init() {
    f := BAfragola阿夫拉戈拉.(*阿夫拉戈拉AfragolaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "afragola",
		TitleName: "阿夫拉戈拉",
		TitleCode: "b_afragola",
	}
}
