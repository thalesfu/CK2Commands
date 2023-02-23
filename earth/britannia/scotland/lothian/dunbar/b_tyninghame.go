package dunbar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 泰宁汉姆TyninghameBarony struct {
	feud.BaseBarony
}

var BTyninghame泰宁汉姆 feud.Barony = &泰宁汉姆TyninghameBarony{}

func init() {
	f := BTyninghame泰宁汉姆.(*泰宁汉姆TyninghameBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "tyninghame",
		TitleName: "泰宁汉姆",
		TitleCode: "b_tyninghame",
	}
}
