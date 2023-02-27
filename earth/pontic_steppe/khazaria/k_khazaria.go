package khazaria

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/aqtobe"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/atyrau"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/bandja"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/itil"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/sakmara"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/khazaria/sarkel"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhazariaKingdom interface {
    feud.Kingdom
    DAqtobe阿克托别() 	aqtobe.AqtobeDuke
    DAtyrau阿特劳() 	atyrau.AtyrauDuke
    DBandja班贾() 	bandja.BandjaDuke
    DItil阿得() 	itil.ItilDuke
    DSakmara萨克马拉() 	sakmara.SakmaraDuke
    DSarkel萨克尔() 	sarkel.SarkelDuke
}

type 可萨利亚KhazariaKingdom struct {
	feud.BaseKingdom
	阿克托别Aqtobe 	aqtobe.AqtobeDuke
	阿特劳Atyrau 	atyrau.AtyrauDuke
	班贾Bandja 	bandja.BandjaDuke
	阿得Itil 	itil.ItilDuke
	萨克马拉Sakmara 	sakmara.SakmaraDuke
	萨克尔Sarkel 	sarkel.SarkelDuke
}

func (k *可萨利亚KhazariaKingdom) DAqtobe阿克托别() aqtobe.AqtobeDuke {
	return k.阿克托别Aqtobe
}
    
func (k *可萨利亚KhazariaKingdom) DAtyrau阿特劳() atyrau.AtyrauDuke {
	return k.阿特劳Atyrau
}
    
func (k *可萨利亚KhazariaKingdom) DBandja班贾() bandja.BandjaDuke {
	return k.班贾Bandja
}
    
func (k *可萨利亚KhazariaKingdom) DItil阿得() itil.ItilDuke {
	return k.阿得Itil
}
    
func (k *可萨利亚KhazariaKingdom) DSakmara萨克马拉() sakmara.SakmaraDuke {
	return k.萨克马拉Sakmara
}
    
func (k *可萨利亚KhazariaKingdom) DSarkel萨克尔() sarkel.SarkelDuke {
	return k.萨克尔Sarkel
}
    
var KKhazaria可萨利亚 KhazariaKingdom = &可萨利亚KhazariaKingdom{}

func init() {
	f := KKhazaria可萨利亚.(*可萨利亚KhazariaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "khazaria",
		TitleName: "可萨利亚",
		TitleCode: "k_khazaria",
		Dukes:  map[string]feud.Duke{},
	}

	f.阿克托别Aqtobe = aqtobe.DAqtobe阿克托别
	f.阿克托别Aqtobe.SetParent(f)
	
	f.阿特劳Atyrau = atyrau.DAtyrau阿特劳
	f.阿特劳Atyrau.SetParent(f)
	
	f.班贾Bandja = bandja.DBandja班贾
	f.班贾Bandja.SetParent(f)
	
	f.阿得Itil = itil.DItil阿得
	f.阿得Itil.SetParent(f)
	
	f.萨克马拉Sakmara = sakmara.DSakmara萨克马拉
	f.萨克马拉Sakmara.SetParent(f)
	
	f.萨克尔Sarkel = sarkel.DSarkel萨克尔
	f.萨克尔Sarkel.SetParent(f)
	
}
