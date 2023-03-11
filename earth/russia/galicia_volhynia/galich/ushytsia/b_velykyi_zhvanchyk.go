package ushytsia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 大日万奇克Velykyi_zhvanchykBarony struct {
	feud.BaseBarony
}

var BVelykyi_zhvanchyk大日万奇克 feud.Barony = &大日万奇克Velykyi_zhvanchykBarony{}

func init() {
    f := BVelykyi_zhvanchyk大日万奇克.(*大日万奇克Velykyi_zhvanchykBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "velykyi_zhvanchyk",
		TitleName: "大日万奇克",
		TitleCode: "b_velykyi_zhvanchyk",
	}
}
