package wolgast

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌瑟多姆UsedomBarony struct {
	feud.BaseBarony
}

var BUsedom乌瑟多姆 feud.Barony = &乌瑟多姆UsedomBarony{}

func init() {
    f := BUsedom乌瑟多姆.(*乌瑟多姆UsedomBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "usedom",
		TitleName: "乌瑟多姆",
		TitleCode: "b_usedom",
	}
}
