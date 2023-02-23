package sinjar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 尼尼微NinevehBarony struct {
	feud.BaseBarony
}

var BNineveh尼尼微 feud.Barony = &尼尼微NinevehBarony{}

func init() {
	f := BNineveh尼尼微.(*尼尼微NinevehBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "nineveh",
		TitleName: "尼尼微",
		TitleCode: "b_nineveh",
	}
}
