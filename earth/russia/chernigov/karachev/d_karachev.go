package karachev

import (
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/karachev/karachev"
	"github.com/thalesfu/CK2Commands/earth/russia/chernigov/karachev/kozelsk"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KarachevDuke interface {
    feud.Duke
    CKarachev卡拉切夫() 	karachev.KarachevCounty
    CKozelsk科泽利斯克() 	kozelsk.KozelskCounty
}

type 卡拉切夫KarachevDuke struct {
	feud.BaseDuke
	卡拉切夫Karachev 	karachev.KarachevCounty
	科泽利斯克Kozelsk 	kozelsk.KozelskCounty
}

func (d *卡拉切夫KarachevDuke) CKarachev卡拉切夫() karachev.KarachevCounty {
	return d.卡拉切夫Karachev
}
    
func (d *卡拉切夫KarachevDuke) CKozelsk科泽利斯克() kozelsk.KozelskCounty {
	return d.科泽利斯克Kozelsk
}
    
var DKarachev卡拉切夫 KarachevDuke = &卡拉切夫KarachevDuke{}

func init() {
	f := DKarachev卡拉切夫.(*卡拉切夫KarachevDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "karachev",
		TitleName: "卡拉切夫",
		TitleCode: "d_karachev",
		Counties:  map[string]feud.County{},
	}

	f.卡拉切夫Karachev = karachev.CKarachev卡拉切夫
	f.卡拉切夫Karachev.SetParent(f)
	
	f.科泽利斯克Kozelsk = kozelsk.CKozelsk科泽利斯克
	f.科泽利斯克Kozelsk.SetParent(f)
	
}
