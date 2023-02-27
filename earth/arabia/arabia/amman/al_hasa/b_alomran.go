package al_hasa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 奥姆兰AlomranBarony struct {
	feud.BaseBarony
}

var BAlomran奥姆兰 feud.Barony = &奥姆兰AlomranBarony{}

func init() {
    f := BAlomran奥姆兰.(*奥姆兰AlomranBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "alomran",
		TitleName: "奥姆兰",
		TitleCode: "b_alomran",
	}
}
