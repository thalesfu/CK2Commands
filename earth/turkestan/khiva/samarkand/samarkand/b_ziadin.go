package samarkand

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 济亚丁ZiadinBarony struct {
	feud.BaseBarony
}

var BZiadin济亚丁 feud.Barony = &济亚丁ZiadinBarony{}

func init() {
	f := BZiadin济亚丁.(*济亚丁ZiadinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ziadin",
		TitleName: "济亚丁",
		TitleCode: "b_ziadin",
	}
}
