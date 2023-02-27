package lubelska

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 文奇纳PodwieprzaBarony struct {
	feud.BaseBarony
}

var BPodwieprza文奇纳 feud.Barony = &文奇纳PodwieprzaBarony{}

func init() {
    f := BPodwieprza文奇纳.(*文奇纳PodwieprzaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "podwieprza",
		TitleName: "文奇纳",
		TitleCode: "b_podwieprza",
	}
}
