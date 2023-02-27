package kharibta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 瑙克拉提斯Naucratis_egyptBarony struct {
	feud.BaseBarony
}

var BNaucratis_egypt瑙克拉提斯 feud.Barony = &瑙克拉提斯Naucratis_egyptBarony{}

func init() {
    f := BNaucratis_egypt瑙克拉提斯.(*瑙克拉提斯Naucratis_egyptBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "naucratis_egypt",
		TitleName: "瑙克拉提斯",
		TitleCode: "b_naucratis_egypt",
	}
}
