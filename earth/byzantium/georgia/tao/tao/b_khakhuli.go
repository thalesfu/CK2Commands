package tao

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 哈胡利KhakhuliBarony struct {
	feud.BaseBarony
}

var BKhakhuli哈胡利 feud.Barony = &哈胡利KhakhuliBarony{}

func init() {
    f := BKhakhuli哈胡利.(*哈胡利KhakhuliBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khakhuli",
		TitleName: "哈胡利",
		TitleCode: "b_khakhuli",
	}
}
