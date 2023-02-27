package kuwait

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 安塞穆夏斯AnthemusiasBarony struct {
	feud.BaseBarony
}

var BAnthemusias安塞穆夏斯 feud.Barony = &安塞穆夏斯AnthemusiasBarony{}

func init() {
    f := BAnthemusias安塞穆夏斯.(*安塞穆夏斯AnthemusiasBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "anthemusias",
		TitleName: "安塞穆夏斯",
		TitleCode: "b_anthemusias",
	}
}
