package or

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿谢利塞AschelisayBarony struct {
	feud.BaseBarony
}

var BAschelisay阿谢利塞 feud.Barony = &阿谢利塞AschelisayBarony{}

func init() {
    f := BAschelisay阿谢利塞.(*阿谢利塞AschelisayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "aschelisay",
		TitleName: "阿谢利塞",
		TitleCode: "b_aschelisay",
	}
}
