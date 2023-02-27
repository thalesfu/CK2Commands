package jiuquan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 陇狄LongdiBarony struct {
	feud.BaseBarony
}

var BLongdi陇狄 feud.Barony = &陇狄LongdiBarony{}

func init() {
    f := BLongdi陇狄.(*陇狄LongdiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "longdi",
		TitleName: "陇狄",
		TitleCode: "b_longdi",
	}
}
