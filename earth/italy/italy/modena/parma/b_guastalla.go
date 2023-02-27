package parma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瓜斯塔拉GuastallaBarony struct {
	feud.BaseBarony
}

var BGuastalla瓜斯塔拉 feud.Barony = &瓜斯塔拉GuastallaBarony{}

func init() {
    f := BGuastalla瓜斯塔拉.(*瓜斯塔拉GuastallaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "guastalla",
		TitleName: "瓜斯塔拉",
		TitleCode: "b_guastalla",
	}
}
