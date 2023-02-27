package chagai

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 迦湿弥罗吉梨Kashmir_kiliBarony struct {
	feud.BaseBarony
}

var BKashmir_kili迦湿弥罗吉梨 feud.Barony = &迦湿弥罗吉梨Kashmir_kiliBarony{}

func init() {
    f := BKashmir_kili迦湿弥罗吉梨.(*迦湿弥罗吉梨Kashmir_kiliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "kashmir_kili",
		TitleName: "迦湿弥罗吉梨",
		TitleCode: "b_kashmir_kili",
	}
}
