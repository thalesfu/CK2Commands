package gurma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 比朗加BilangaBarony struct {
	feud.BaseBarony
}

var BBilanga比朗加 feud.Barony = &比朗加BilangaBarony{}

func init() {
	f := BBilanga比朗加.(*比朗加BilangaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "bilanga",
		TitleName: "比朗加",
		TitleCode: "b_bilanga",
	}
}
