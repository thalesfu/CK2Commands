package foix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 蒙塞居尔MontsegurBarony struct {
	feud.BaseBarony
}

var BMontsegur蒙塞居尔 feud.Barony = &蒙塞居尔MontsegurBarony{}

func init() {
	f := BMontsegur蒙塞居尔.(*蒙塞居尔MontsegurBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "montsegur",
		TitleName: "蒙塞居尔",
		TitleCode: "b_montsegur",
	}
}
