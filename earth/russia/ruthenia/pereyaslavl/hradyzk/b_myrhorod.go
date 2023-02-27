package hradyzk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 米尔哥罗德MyrhorodBarony struct {
	feud.BaseBarony
}

var BMyrhorod米尔哥罗德 feud.Barony = &米尔哥罗德MyrhorodBarony{}

func init() {
    f := BMyrhorod米尔哥罗德.(*米尔哥罗德MyrhorodBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "myrhorod",
		TitleName: "米尔哥罗德",
		TitleCode: "b_myrhorod",
	}
}
