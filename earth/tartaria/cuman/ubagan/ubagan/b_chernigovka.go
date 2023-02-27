package ubagan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 切尔尼戈夫卡ChernigovkaBarony struct {
	feud.BaseBarony
}

var BChernigovka切尔尼戈夫卡 feud.Barony = &切尔尼戈夫卡ChernigovkaBarony{}

func init() {
    f := BChernigovka切尔尼戈夫卡.(*切尔尼戈夫卡ChernigovkaBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "chernigovka",
		TitleName: "切尔尼戈夫卡",
		TitleCode: "b_chernigovka",
	}
}
