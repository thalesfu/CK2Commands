package konjikala

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌卢格_杰佩UlugdepeBarony struct {
	feud.BaseBarony
}

var BUlugdepe乌卢格_杰佩 feud.Barony = &乌卢格_杰佩UlugdepeBarony{}

func init() {
    f := BUlugdepe乌卢格_杰佩.(*乌卢格_杰佩UlugdepeBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulugdepe",
		TitleName: "乌卢格_杰佩",
		TitleCode: "b_ulugdepe",
	}
}
