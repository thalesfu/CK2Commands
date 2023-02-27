package vagay

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓦盖VagayBarony struct {
	feud.BaseBarony
}

var BVagay瓦盖 feud.Barony = &瓦盖VagayBarony{}

func init() {
    f := BVagay瓦盖.(*瓦盖VagayBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "vagay",
		TitleName: "瓦盖",
		TitleCode: "b_vagay",
	}
}
