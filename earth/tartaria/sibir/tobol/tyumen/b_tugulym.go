package tyumen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图古雷姆TugulymBarony struct {
	feud.BaseBarony
}

var BTugulym图古雷姆 feud.Barony = &图古雷姆TugulymBarony{}

func init() {
    f := BTugulym图古雷姆.(*图古雷姆TugulymBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tugulym",
		TitleName: "图古雷姆",
		TitleCode: "b_tugulym",
	}
}
