package rummah

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 喀斯苏曼QaisumahBarony struct {
	feud.BaseBarony
}

var BQaisumah喀斯苏曼 feud.Barony = &喀斯苏曼QaisumahBarony{}

func init() {
	f := BQaisumah喀斯苏曼.(*喀斯苏曼QaisumahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qaisumah",
		TitleName: "喀斯苏曼",
		TitleCode: "b_qaisumah",
	}
}
