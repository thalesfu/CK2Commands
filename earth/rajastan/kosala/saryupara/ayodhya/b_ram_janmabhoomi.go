package ayodhya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 罗摩降诞处Ram_janmabhoomiBarony struct {
	feud.BaseBarony
}

var BRam_janmabhoomi罗摩降诞处 feud.Barony = &罗摩降诞处Ram_janmabhoomiBarony{}

func init() {
    f := BRam_janmabhoomi罗摩降诞处.(*罗摩降诞处Ram_janmabhoomiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ram_janmabhoomi",
		TitleName: "罗摩降诞处",
		TitleCode: "b_ram_janmabhoomi",
	}
}
