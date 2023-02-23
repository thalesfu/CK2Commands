package amastris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 克罗姆纳CromnaBarony struct {
	feud.BaseBarony
}

var BCromna克罗姆纳 feud.Barony = &克罗姆纳CromnaBarony{}

func init() {
	f := BCromna克罗姆纳.(*克罗姆纳CromnaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "cromna",
		TitleName: "克罗姆纳",
		TitleCode: "b_cromna",
	}
}
