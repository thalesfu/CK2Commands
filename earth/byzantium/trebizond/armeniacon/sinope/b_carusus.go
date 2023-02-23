package sinope

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡洛索斯CarususBarony struct {
	feud.BaseBarony
}

var BCarusus卡洛索斯 feud.Barony = &卡洛索斯CarususBarony{}

func init() {
	f := BCarusus卡洛索斯.(*卡洛索斯CarususBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "carusus",
		TitleName: "卡洛索斯",
		TitleCode: "b_carusus",
	}
}
