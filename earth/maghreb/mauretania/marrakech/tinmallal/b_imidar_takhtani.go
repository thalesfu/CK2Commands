package tinmallal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 伊米德尔泰赫泰尼Imidar_takhtaniBarony struct {
	feud.BaseBarony
}

var BImidar_takhtani伊米德尔泰赫泰尼 feud.Barony = &伊米德尔泰赫泰尼Imidar_takhtaniBarony{}

func init() {
    f := BImidar_takhtani伊米德尔泰赫泰尼.(*伊米德尔泰赫泰尼Imidar_takhtaniBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "imidar_takhtani",
		TitleName: "伊米德尔泰赫泰尼",
		TitleCode: "b_imidar_takhtani",
	}
}
