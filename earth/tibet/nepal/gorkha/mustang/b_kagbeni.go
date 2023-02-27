package mustang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡贝尼KagbeniBarony struct {
	feud.BaseBarony
}

var BKagbeni卡贝尼 feud.Barony = &卡贝尼KagbeniBarony{}

func init() {
    f := BKagbeni卡贝尼.(*卡贝尼KagbeniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kagbeni",
		TitleName: "卡贝尼",
		TitleCode: "b_kagbeni",
	}
}
