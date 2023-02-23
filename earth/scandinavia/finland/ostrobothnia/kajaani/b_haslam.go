package kajaani

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈斯拉姆HaslamBarony struct {
	feud.BaseBarony
}

var BHaslam哈斯拉姆 feud.Barony = &哈斯拉姆HaslamBarony{}

func init() {
	f := BHaslam哈斯拉姆.(*哈斯拉姆HaslamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "haslam",
		TitleName: "哈斯拉姆",
		TitleCode: "b_haslam",
	}
}
