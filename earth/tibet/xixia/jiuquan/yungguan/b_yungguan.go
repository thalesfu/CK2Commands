package yungguan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阳关YungguanBarony struct {
	feud.BaseBarony
}

var BYungguan阳关 feud.Barony = &阳关YungguanBarony{}

func init() {
    f := BYungguan阳关.(*阳关YungguanBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yungguan",
		TitleName: "阳关",
		TitleCode: "b_yungguan",
	}
}
