package naro-fominsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 库宾卡KubinkaBarony struct {
	feud.BaseBarony
}

var BKubinka库宾卡 feud.Barony = &库宾卡KubinkaBarony{}

func init() {
    f := BKubinka库宾卡.(*库宾卡KubinkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kubinka",
		TitleName: "库宾卡",
		TitleCode: "b_kubinka",
	}
}
