package kagha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊科特奥比奥恩苏Ikot_obio_nsuBarony struct {
	feud.BaseBarony
}

var BIkot_obio_nsu伊科特奥比奥恩苏 feud.Barony = &伊科特奥比奥恩苏Ikot_obio_nsuBarony{}

func init() {
    f := BIkot_obio_nsu伊科特奥比奥恩苏.(*伊科特奥比奥恩苏Ikot_obio_nsuBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ikot_obio_nsu",
		TitleName: "伊科特奥比奥恩苏",
		TitleCode: "b_ikot_obio_nsu",
	}
}
