package az_zarqa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 乌姆胡尔盖KhurqahBarony struct {
	feud.BaseBarony
}

var BKhurqah乌姆胡尔盖 feud.Barony = &乌姆胡尔盖KhurqahBarony{}

func init() {
    f := BKhurqah乌姆胡尔盖.(*乌姆胡尔盖KhurqahBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "khurqah",
		TitleName: "乌姆胡尔盖",
		TitleCode: "b_khurqah",
	}
}
