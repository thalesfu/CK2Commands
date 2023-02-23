package belgorod

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克乌谢尼CauseniBarony struct {
	feud.BaseBarony
}

var BCauseni克乌谢尼 feud.Barony = &克乌谢尼CauseniBarony{}

func init() {
	f := BCauseni克乌谢尼.(*克乌谢尼CauseniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "causeni",
		TitleName: "克乌谢尼",
		TitleCode: "b_causeni",
	}
}
