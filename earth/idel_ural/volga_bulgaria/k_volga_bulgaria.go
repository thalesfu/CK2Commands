package volga_bulgaria

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/bulgar"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/cheremisa"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/kazan"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/maris"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/volga_bulgaria/mordvins"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Volga_bulgariaKingdom interface {
    feud.Kingdom
    DBulgar保加尔() 	bulgar.BulgarDuke
    DCheremisa切列米莎() 	cheremisa.CheremisaDuke
    DKazan喀山() 	kazan.KazanDuke
    DMaris马里() 	maris.MarisDuke
    DMordvins莫尔多瓦() 	mordvins.MordvinsDuke
}

type 伏尔加保加利亚Volga_bulgariaKingdom struct {
	feud.BaseKingdom
	保加尔Bulgar 	bulgar.BulgarDuke
	切列米莎Cheremisa 	cheremisa.CheremisaDuke
	喀山Kazan 	kazan.KazanDuke
	马里Maris 	maris.MarisDuke
	莫尔多瓦Mordvins 	mordvins.MordvinsDuke
}

func (k *伏尔加保加利亚Volga_bulgariaKingdom) DBulgar保加尔() bulgar.BulgarDuke {
	return k.保加尔Bulgar
}
    
func (k *伏尔加保加利亚Volga_bulgariaKingdom) DCheremisa切列米莎() cheremisa.CheremisaDuke {
	return k.切列米莎Cheremisa
}
    
func (k *伏尔加保加利亚Volga_bulgariaKingdom) DKazan喀山() kazan.KazanDuke {
	return k.喀山Kazan
}
    
func (k *伏尔加保加利亚Volga_bulgariaKingdom) DMaris马里() maris.MarisDuke {
	return k.马里Maris
}
    
func (k *伏尔加保加利亚Volga_bulgariaKingdom) DMordvins莫尔多瓦() mordvins.MordvinsDuke {
	return k.莫尔多瓦Mordvins
}
    
var KVolga_bulgaria伏尔加保加利亚 Volga_bulgariaKingdom = &伏尔加保加利亚Volga_bulgariaKingdom{}

func init() {
	f := KVolga_bulgaria伏尔加保加利亚.(*伏尔加保加利亚Volga_bulgariaKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "volga_bulgaria",
		TitleName: "伏尔加保加利亚",
		TitleCode: "k_volga_bulgaria",
		Dukes:  map[string]feud.Duke{},
	}

	f.保加尔Bulgar = bulgar.DBulgar保加尔
	f.保加尔Bulgar.SetParent(f)
	
	f.切列米莎Cheremisa = cheremisa.DCheremisa切列米莎
	f.切列米莎Cheremisa.SetParent(f)
	
	f.喀山Kazan = kazan.DKazan喀山
	f.喀山Kazan.SetParent(f)
	
	f.马里Maris = maris.DMaris马里
	f.马里Maris.SetParent(f)
	
	f.莫尔多瓦Mordvins = mordvins.DMordvins莫尔多瓦
	f.莫尔多瓦Mordvins.SetParent(f)
	
}
