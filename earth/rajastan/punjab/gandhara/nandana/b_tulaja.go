package nandana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陀罗奢TulajaBarony struct {
	feud.BaseBarony
}

var BTulaja陀罗奢 feud.Barony = &陀罗奢TulajaBarony{}

func init() {
    f := BTulaja陀罗奢.(*陀罗奢TulajaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tulaja",
		TitleName: "陀罗奢",
		TitleCode: "b_tulaja",
	}
}
