package northumberland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尼克AlnwickBarony struct {
	feud.BaseBarony
}

var BAlnwick阿尼克 feud.Barony = &阿尼克AlnwickBarony{}

func init() {
    f := BAlnwick阿尼克.(*阿尼克AlnwickBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alnwick",
		TitleName: "阿尼克",
		TitleCode: "b_alnwick",
	}
}
