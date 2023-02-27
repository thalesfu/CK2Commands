package novgorod_seversky

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 阿尔特诺夫卡AltynivkaBarony struct {
	feud.BaseBarony
}

var BAltynivka阿尔特诺夫卡 feud.Barony = &阿尔特诺夫卡AltynivkaBarony{}

func init() {
    f := BAltynivka阿尔特诺夫卡.(*阿尔特诺夫卡AltynivkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "altynivka",
		TitleName: "阿尔特诺夫卡",
		TitleCode: "b_altynivka",
	}
}
