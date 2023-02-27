package abauj

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type 沙托劳尔尧乌伊海伊SatoraljaujhelyBarony struct {
	feud.BaseBarony
}

var BSatoraljaujhely沙托劳尔尧乌伊海伊 feud.Barony = &沙托劳尔尧乌伊海伊SatoraljaujhelyBarony{}

func init() {
    f := BSatoraljaujhely沙托劳尔尧乌伊海伊.(*沙托劳尔尧乌伊海伊SatoraljaujhelyBarony)
	f.BaseBarony = feud.BaseBarony{
		Title:     "satoraljaujhely",
		TitleName: "沙托劳尔尧乌伊海伊",
		TitleCode: "b_satoraljaujhely",
	}
}
