package holland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿姆斯特丹AmsterdamBarony struct {
	feud.BaseBarony
}

var BAmsterdam阿姆斯特丹 feud.Barony = &阿姆斯特丹AmsterdamBarony{}

func init() {
	f := BAmsterdam阿姆斯特丹.(*阿姆斯特丹AmsterdamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "amsterdam",
		TitleName: "阿姆斯特丹",
		TitleCode: "b_amsterdam",
	}
}
