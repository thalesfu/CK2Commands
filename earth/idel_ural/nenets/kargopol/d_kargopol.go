package kargopol

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/kargopol/kargopol"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/kargopol/romny"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/nenets/kargopol/trans_portage"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KargopolDuke interface {
    feud.Duke
    CKargopol卡尔戈波尔() 	kargopol.KargopolCounty
    CRomny罗姆内() 	romny.RomnyCounty
    CTrans_portage奥涅加() 	trans_portage.Trans_portageCounty
}

type 卡尔戈波尔KargopolDuke struct {
	feud.BaseDuke
	卡尔戈波尔Kargopol 	kargopol.KargopolCounty
	罗姆内Romny 	romny.RomnyCounty
	奥涅加Trans_portage 	trans_portage.Trans_portageCounty
}

func (d *卡尔戈波尔KargopolDuke) CKargopol卡尔戈波尔() kargopol.KargopolCounty {
	return d.卡尔戈波尔Kargopol
}
    
func (d *卡尔戈波尔KargopolDuke) CRomny罗姆内() romny.RomnyCounty {
	return d.罗姆内Romny
}
    
func (d *卡尔戈波尔KargopolDuke) CTrans_portage奥涅加() trans_portage.Trans_portageCounty {
	return d.奥涅加Trans_portage
}
    
var DKargopol卡尔戈波尔 KargopolDuke = &卡尔戈波尔KargopolDuke{}

func init() {
	f := DKargopol卡尔戈波尔.(*卡尔戈波尔KargopolDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kargopol",
		TitleName: "卡尔戈波尔",
		TitleCode: "d_kargopol",
		Counties:  map[string]feud.County{},
	}

	f.卡尔戈波尔Kargopol = kargopol.CKargopol卡尔戈波尔
	f.卡尔戈波尔Kargopol.SetParent(f)
	
	f.罗姆内Romny = romny.CRomny罗姆内
	f.罗姆内Romny.SetParent(f)
	
	f.奥涅加Trans_portage = trans_portage.CTrans_portage奥涅加
	f.奥涅加Trans_portage.SetParent(f)
	
}
