package mecklemburg

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 格雷沃斯米伦GrevesmuhlenBarony struct {
	feud.BaseBarony
}

var BGrevesmuhlen格雷沃斯米伦 feud.Barony = &格雷沃斯米伦GrevesmuhlenBarony{}

func init() {
    f := BGrevesmuhlen格雷沃斯米伦.(*格雷沃斯米伦GrevesmuhlenBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "grevesmuhlen",
		TitleName: "格雷沃斯米伦",
		TitleCode: "b_grevesmuhlen",
	}
}
