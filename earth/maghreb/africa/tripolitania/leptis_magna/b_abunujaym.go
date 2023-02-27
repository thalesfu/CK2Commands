package leptis_magna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿布努贾AbunujaymBarony struct {
	feud.BaseBarony
}

var BAbunujaym阿布努贾 feud.Barony = &阿布努贾AbunujaymBarony{}

func init() {
    f := BAbunujaym阿布努贾.(*阿布努贾AbunujaymBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "abunujaym",
		TitleName: "阿布努贾",
		TitleCode: "b_abunujaym",
	}
}
