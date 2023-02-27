package viviers

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 拉让蒂耶尔LargentiereBarony struct {
	feud.BaseBarony
}

var BLargentiere拉让蒂耶尔 feud.Barony = &拉让蒂耶尔LargentiereBarony{}

func init() {
    f := BLargentiere拉让蒂耶尔.(*拉让蒂耶尔LargentiereBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "largentiere",
		TitleName: "拉让蒂耶尔",
		TitleCode: "b_largentiere",
	}
}
