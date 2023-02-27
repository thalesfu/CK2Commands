package orava

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 图尔多欣TurdossinBarony struct {
	feud.BaseBarony
}

var BTurdossin图尔多欣 feud.Barony = &图尔多欣TurdossinBarony{}

func init() {
    f := BTurdossin图尔多欣.(*图尔多欣TurdossinBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "turdossin",
		TitleName: "图尔多欣",
		TitleCode: "b_turdossin",
	}
}
