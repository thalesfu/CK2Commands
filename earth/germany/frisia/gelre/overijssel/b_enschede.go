package overijssel

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 恩斯赫德EnschedeBarony struct {
	feud.BaseBarony
}

var BEnschede恩斯赫德 feud.Barony = &恩斯赫德EnschedeBarony{}

func init() {
    f := BEnschede恩斯赫德.(*恩斯赫德EnschedeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "enschede",
		TitleName: "恩斯赫德",
		TitleCode: "b_enschede",
	}
}
