package udine

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 塔尔琴托TarcentoBarony struct {
	feud.BaseBarony
}

var BTarcento塔尔琴托 feud.Barony = &塔尔琴托TarcentoBarony{}

func init() {
    f := BTarcento塔尔琴托.(*塔尔琴托TarcentoBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tarcento",
		TitleName: "塔尔琴托",
		TitleCode: "b_tarcento",
	}
}
