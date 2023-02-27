package kudymkar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 谢伊瓦SeyvaBarony struct {
	feud.BaseBarony
}

var BSeyva谢伊瓦 feud.Barony = &谢伊瓦SeyvaBarony{}

func init() {
    f := BSeyva谢伊瓦.(*谢伊瓦SeyvaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "seyva",
		TitleName: "谢伊瓦",
		TitleCode: "b_seyva",
	}
}
