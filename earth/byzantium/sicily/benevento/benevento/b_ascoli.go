package benevento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿斯科利AscoliBarony struct {
	feud.BaseBarony
}

var BAscoli阿斯科利 feud.Barony = &阿斯科利AscoliBarony{}

func init() {
    f := BAscoli阿斯科利.(*阿斯科利AscoliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ascoli",
		TitleName: "阿斯科利",
		TitleCode: "b_ascoli",
	}
}
