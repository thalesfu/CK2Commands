package damascus

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/damascus/al_mafraq"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/damascus/amman"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/damascus/az_zarqa"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/damascus/damascus"
	"github.com/thalesfu/CK2Commands/earth/arabia/syria/damascus/irbid"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DamascusDuke interface {
    feud.Duke
    CAl_mafraq马弗拉克() 	al_mafraq.Al_mafraqCounty
    CAmman安曼() 	amman.AmmanCounty
    CAz_zarqa扎尔卡() 	az_zarqa.Az_zarqaCounty
    CDamascus大马士革() 	damascus.DamascusCounty
    CIrbid伊尔比德() 	irbid.IrbidCounty
}

type 大马士革DamascusDuke struct {
	feud.BaseDuke
	马弗拉克Al_mafraq 	al_mafraq.Al_mafraqCounty
	安曼Amman 	amman.AmmanCounty
	扎尔卡Az_zarqa 	az_zarqa.Az_zarqaCounty
	大马士革Damascus 	damascus.DamascusCounty
	伊尔比德Irbid 	irbid.IrbidCounty
}

func (d *大马士革DamascusDuke) CAl_mafraq马弗拉克() al_mafraq.Al_mafraqCounty {
	return d.马弗拉克Al_mafraq
}
    
func (d *大马士革DamascusDuke) CAmman安曼() amman.AmmanCounty {
	return d.安曼Amman
}
    
func (d *大马士革DamascusDuke) CAz_zarqa扎尔卡() az_zarqa.Az_zarqaCounty {
	return d.扎尔卡Az_zarqa
}
    
func (d *大马士革DamascusDuke) CDamascus大马士革() damascus.DamascusCounty {
	return d.大马士革Damascus
}
    
func (d *大马士革DamascusDuke) CIrbid伊尔比德() irbid.IrbidCounty {
	return d.伊尔比德Irbid
}
    
var DDamascus大马士革 DamascusDuke = &大马士革DamascusDuke{}

func init() {
	f := DDamascus大马士革.(*大马士革DamascusDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "damascus",
		TitleName: "大马士革",
		TitleCode: "d_damascus",
		Counties:  map[string]feud.County{},
	}

	f.马弗拉克Al_mafraq = al_mafraq.CAl_mafraq马弗拉克
	f.马弗拉克Al_mafraq.SetParent(f)
	
	f.安曼Amman = amman.CAmman安曼
	f.安曼Amman.SetParent(f)
	
	f.扎尔卡Az_zarqa = az_zarqa.CAz_zarqa扎尔卡
	f.扎尔卡Az_zarqa.SetParent(f)
	
	f.大马士革Damascus = damascus.CDamascus大马士革
	f.大马士革Damascus.SetParent(f)
	
	f.伊尔比德Irbid = irbid.CIrbid伊尔比德
	f.伊尔比德Irbid.SetParent(f)
	
}
