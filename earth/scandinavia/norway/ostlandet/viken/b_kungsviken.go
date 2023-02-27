package viken

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 孔斯维肯KungsvikenBarony struct {
	feud.BaseBarony
}

var BKungsviken孔斯维肯 feud.Barony = &孔斯维肯KungsvikenBarony{}

func init() {
    f := BKungsviken孔斯维肯.(*孔斯维肯KungsvikenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kungsviken",
		TitleName: "孔斯维肯",
		TitleCode: "b_kungsviken",
	}
}
