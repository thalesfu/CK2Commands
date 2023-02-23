package bambuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 亚雷斯纳YaresnaBarony struct {
	feud.BaseBarony
}

var BYaresna亚雷斯纳 feud.Barony = &亚雷斯纳YaresnaBarony{}

func init() {
	f := BYaresna亚雷斯纳.(*亚雷斯纳YaresnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "yaresna",
		TitleName: "亚雷斯纳",
		TitleCode: "b_yaresna",
	}
}
