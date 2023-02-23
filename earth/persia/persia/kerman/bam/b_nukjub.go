package bam

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 努克祖NukjubBarony struct {
	feud.BaseBarony
}

var BNukjub努克祖 feud.Barony = &努克祖NukjubBarony{}

func init() {
	f := BNukjub努克祖.(*努克祖NukjubBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nukjub",
		TitleName: "努克祖",
		TitleCode: "b_nukjub",
	}
}
