package schwyz

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格拉鲁斯GlarusBarony struct {
	feud.BaseBarony
}

var BGlarus格拉鲁斯 feud.Barony = &格拉鲁斯GlarusBarony{}

func init() {
	f := BGlarus格拉鲁斯.(*格拉鲁斯GlarusBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "glarus",
		TitleName: "格拉鲁斯",
		TitleCode: "b_glarus",
	}
}
