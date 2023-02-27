package zavarot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库亚KuyaBarony struct {
	feud.BaseBarony
}

var BKuya库亚 feud.Barony = &库亚KuyaBarony{}

func init() {
    f := BKuya库亚.(*库亚KuyaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kuya",
		TitleName: "库亚",
		TitleCode: "b_kuya",
	}
}
