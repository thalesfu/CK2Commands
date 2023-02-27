package holland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈勒姆HaarlemBarony struct {
	feud.BaseBarony
}

var BHaarlem哈勒姆 feud.Barony = &哈勒姆HaarlemBarony{}

func init() {
    f := BHaarlem哈勒姆.(*哈勒姆HaarlemBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haarlem",
		TitleName: "哈勒姆",
		TitleCode: "b_haarlem",
	}
}
