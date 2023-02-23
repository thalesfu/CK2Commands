package volkovysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 利亚哈维奇LyakhavichyBarony struct {
	feud.BaseBarony
}

var BLyakhavichy利亚哈维奇 feud.Barony = &利亚哈维奇LyakhavichyBarony{}

func init() {
	f := BLyakhavichy利亚哈维奇.(*利亚哈维奇LyakhavichyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "lyakhavichy",
		TitleName: "利亚哈维奇",
		TitleCode: "b_lyakhavichy",
	}
}
