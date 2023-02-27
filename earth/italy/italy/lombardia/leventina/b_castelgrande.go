package leventina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 卡斯特尔格兰德CastelgrandeBarony struct {
	feud.BaseBarony
}

var BCastelgrande卡斯特尔格兰德 feud.Barony = &卡斯特尔格兰德CastelgrandeBarony{}

func init() {
    f := BCastelgrande卡斯特尔格兰德.(*卡斯特尔格兰德CastelgrandeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "castelgrande",
		TitleName: "卡斯特尔格兰德",
		TitleCode: "b_castelgrande",
	}
}
