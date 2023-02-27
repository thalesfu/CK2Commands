package vairagara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 贡迪亚GondiaBarony struct {
	feud.BaseBarony
}

var BGondia贡迪亚 feud.Barony = &贡迪亚GondiaBarony{}

func init() {
    f := BGondia贡迪亚.(*贡迪亚GondiaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "gondia",
		TitleName: "贡迪亚",
		TitleCode: "b_gondia",
	}
}
