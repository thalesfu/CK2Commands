package oleshye

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼古拉耶夫MykolaivBarony struct {
	feud.BaseBarony
}

var BMykolaiv尼古拉耶夫 feud.Barony = &尼古拉耶夫MykolaivBarony{}

func init() {
    f := BMykolaiv尼古拉耶夫.(*尼古拉耶夫MykolaivBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "mykolaiv",
		TitleName: "尼古拉耶夫",
		TitleCode: "b_mykolaiv",
	}
}
