package navasarika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 达兰补罗DharampurBarony struct {
	feud.BaseBarony
}

var BDharampur达兰补罗 feud.Barony = &达兰补罗DharampurBarony{}

func init() {
	f := BDharampur达兰补罗.(*达兰补罗DharampurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "dharampur",
		TitleName: "达兰补罗",
		TitleCode: "b_dharampur",
	}
}
