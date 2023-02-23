package herjedalen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌尔夫谢拉UlvkallaBarony struct {
	feud.BaseBarony
}

var BUlvkalla乌尔夫谢拉 feud.Barony = &乌尔夫谢拉UlvkallaBarony{}

func init() {
	f := BUlvkalla乌尔夫谢拉.(*乌尔夫谢拉UlvkallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "ulvkalla",
		TitleName: "乌尔夫谢拉",
		TitleCode: "b_ulvkalla",
	}
}
