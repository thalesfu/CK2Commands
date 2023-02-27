package artux

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 波斯喀木PoskamBarony struct {
	feud.BaseBarony
}

var BPoskam波斯喀木 feud.Barony = &波斯喀木PoskamBarony{}

func init() {
    f := BPoskam波斯喀木.(*波斯喀木PoskamBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "poskam",
		TitleName: "波斯喀木",
		TitleCode: "b_poskam",
	}
}
