package qayaliq

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 海押立QayaliqBarony struct {
	feud.BaseBarony
}

var BQayaliq海押立 feud.Barony = &海押立QayaliqBarony{}

func init() {
    f := BQayaliq海押立.(*海押立QayaliqBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "qayaliq",
		TitleName: "海押立",
		TitleCode: "b_qayaliq",
	}
}
