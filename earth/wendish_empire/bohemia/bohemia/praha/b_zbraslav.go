package praha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 兹布拉斯拉夫ZbraslavBarony struct {
	feud.BaseBarony
}

var BZbraslav兹布拉斯拉夫 feud.Barony = &兹布拉斯拉夫ZbraslavBarony{}

func init() {
    f := BZbraslav兹布拉斯拉夫.(*兹布拉斯拉夫ZbraslavBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "zbraslav",
		TitleName: "兹布拉斯拉夫",
		TitleCode: "b_zbraslav",
	}
}
