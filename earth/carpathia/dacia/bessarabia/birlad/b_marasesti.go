package birlad

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 默勒谢什蒂MarasestiBarony struct {
	feud.BaseBarony
}

var BMarasesti默勒谢什蒂 feud.Barony = &默勒谢什蒂MarasestiBarony{}

func init() {
	f := BMarasesti默勒谢什蒂.(*默勒谢什蒂MarasestiBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "marasesti",
		TitleName: "默勒谢什蒂",
		TitleCode: "b_marasesti",
	}
}
