package nabadwipa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 般遮剌那PancharatnaBarony struct {
	feud.BaseBarony
}

var BPancharatna般遮剌那 feud.Barony = &般遮剌那PancharatnaBarony{}

func init() {
    f := BPancharatna般遮剌那.(*般遮剌那PancharatnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "pancharatna",
		TitleName: "般遮剌那",
		TitleCode: "b_pancharatna",
	}
}
