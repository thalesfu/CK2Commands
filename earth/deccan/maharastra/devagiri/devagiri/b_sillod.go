package devagiri

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 锡洛德SillodBarony struct {
	feud.BaseBarony
}

var BSillod锡洛德 feud.Barony = &锡洛德SillodBarony{}

func init() {
	f := BSillod锡洛德.(*锡洛德SillodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "sillod",
		TitleName: "锡洛德",
		TitleCode: "b_sillod",
	}
}
