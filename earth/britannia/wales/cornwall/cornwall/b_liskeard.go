package cornwall

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利斯卡德LiskeardBarony struct {
	feud.BaseBarony
}

var BLiskeard利斯卡德 feud.Barony = &利斯卡德LiskeardBarony{}

func init() {
    f := BLiskeard利斯卡德.(*利斯卡德LiskeardBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "liskeard",
		TitleName: "利斯卡德",
		TitleCode: "b_liskeard",
	}
}
