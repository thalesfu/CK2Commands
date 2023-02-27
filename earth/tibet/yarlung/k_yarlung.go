package yarlung

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/bhutan"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/lhasa"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/nagchu"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/shigatse"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/sumparu"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/yarlung"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YarlungKingdom interface {
    feud.Kingdom
    DBhutan不丹() 	bhutan.BhutanDuke
    DLhasa逻些() 	lhasa.LhasaDuke
    DNagchu那曲() 	nagchu.NagchuDuke
    DShigatse日喀则() 	shigatse.ShigatseDuke
    DSumparu孙波如() 	sumparu.SumparuDuke
    DYarlung雅砻() 	yarlung.YarlungDuke
}

type 卫藏YarlungKingdom struct {
	feud.BaseKingdom
	不丹Bhutan 	bhutan.BhutanDuke
	逻些Lhasa 	lhasa.LhasaDuke
	那曲Nagchu 	nagchu.NagchuDuke
	日喀则Shigatse 	shigatse.ShigatseDuke
	孙波如Sumparu 	sumparu.SumparuDuke
	雅砻Yarlung 	yarlung.YarlungDuke
}

func (k *卫藏YarlungKingdom) DBhutan不丹() bhutan.BhutanDuke {
	return k.不丹Bhutan
}
    
func (k *卫藏YarlungKingdom) DLhasa逻些() lhasa.LhasaDuke {
	return k.逻些Lhasa
}
    
func (k *卫藏YarlungKingdom) DNagchu那曲() nagchu.NagchuDuke {
	return k.那曲Nagchu
}
    
func (k *卫藏YarlungKingdom) DShigatse日喀则() shigatse.ShigatseDuke {
	return k.日喀则Shigatse
}
    
func (k *卫藏YarlungKingdom) DSumparu孙波如() sumparu.SumparuDuke {
	return k.孙波如Sumparu
}
    
func (k *卫藏YarlungKingdom) DYarlung雅砻() yarlung.YarlungDuke {
	return k.雅砻Yarlung
}
    
var KYarlung卫藏 YarlungKingdom = &卫藏YarlungKingdom{}

func init() {
	f := KYarlung卫藏.(*卫藏YarlungKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "yarlung",
		TitleName: "卫藏",
		TitleCode: "k_yarlung",
		Dukes:  map[string]feud.Duke{},
	}

	f.不丹Bhutan = bhutan.DBhutan不丹
	f.不丹Bhutan.SetParent(f)
	
	f.逻些Lhasa = lhasa.DLhasa逻些
	f.逻些Lhasa.SetParent(f)
	
	f.那曲Nagchu = nagchu.DNagchu那曲
	f.那曲Nagchu.SetParent(f)
	
	f.日喀则Shigatse = shigatse.DShigatse日喀则
	f.日喀则Shigatse.SetParent(f)
	
	f.孙波如Sumparu = sumparu.DSumparu孙波如
	f.孙波如Sumparu.SetParent(f)
	
	f.雅砻Yarlung = yarlung.DYarlung雅砻
	f.雅砻Yarlung.SetParent(f)
	
}
