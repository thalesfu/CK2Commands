package altay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阔克阿尕什KokagaxBarony struct {
	feud.BaseBarony
}

var BKokagax阔克阿尕什 feud.Barony = &阔克阿尕什KokagaxBarony{}

func init() {
    f := BKokagax阔克阿尕什.(*阔克阿尕什KokagaxBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kokagax",
		TitleName: "阔克阿尕什",
		TitleCode: "b_kokagax",
	}
}
