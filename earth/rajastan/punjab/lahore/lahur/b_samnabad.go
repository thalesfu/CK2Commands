package lahur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 萨姆那巴德SamnabadBarony struct {
	feud.BaseBarony
}

var BSamnabad萨姆那巴德 feud.Barony = &萨姆那巴德SamnabadBarony{}

func init() {
	f := BSamnabad萨姆那巴德.(*萨姆那巴德SamnabadBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "samnabad",
		TitleName: "萨姆那巴德",
		TitleCode: "b_samnabad",
	}
}
