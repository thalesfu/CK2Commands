package iasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IasiCounty interface {
    feud.County
    BCernauti切尔讷乌齐() 	feud.Barony
    BDorohoi多罗霍伊() 	feud.Barony
    BHarlau赫尔勒乌() 	feud.Barony
    BHotin霍京() 	feud.Barony
    BIasi雅西() 	feud.Barony
    BUngheni温盖尼() 	feud.Barony
    BVarzaresti沃尔泽雷什蒂() 	feud.Barony
}

type 雅西IasiCounty struct {
	feud.BaseCounty
	切尔讷乌齐Cernauti 	feud.Barony
	多罗霍伊Dorohoi 	feud.Barony
	赫尔勒乌Harlau 	feud.Barony
	霍京Hotin 	feud.Barony
	雅西Iasi 	feud.Barony
	温盖尼Ungheni 	feud.Barony
	沃尔泽雷什蒂Varzaresti 	feud.Barony
}

func (c *雅西IasiCounty) BCernauti切尔讷乌齐() feud.Barony {
	return c.切尔讷乌齐Cernauti
}
    
func (c *雅西IasiCounty) BDorohoi多罗霍伊() feud.Barony {
	return c.多罗霍伊Dorohoi
}
    
func (c *雅西IasiCounty) BHarlau赫尔勒乌() feud.Barony {
	return c.赫尔勒乌Harlau
}
    
func (c *雅西IasiCounty) BHotin霍京() feud.Barony {
	return c.霍京Hotin
}
    
func (c *雅西IasiCounty) BIasi雅西() feud.Barony {
	return c.雅西Iasi
}
    
func (c *雅西IasiCounty) BUngheni温盖尼() feud.Barony {
	return c.温盖尼Ungheni
}
    
func (c *雅西IasiCounty) BVarzaresti沃尔泽雷什蒂() feud.Barony {
	return c.沃尔泽雷什蒂Varzaresti
}
    
var CIasi雅西 IasiCounty = &雅西IasiCounty{}

func init() {
	f := CIasi雅西.(*雅西IasiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1638",
		Title:     "iasi",
		TitleName: "雅西",
		TitleCode: "c_iasi",
		Baronies:  map[string]feud.Barony{},
	}

	f.切尔讷乌齐Cernauti = BCernauti切尔讷乌齐
	f.切尔讷乌齐Cernauti.SetParent(f)
	
	f.多罗霍伊Dorohoi = BDorohoi多罗霍伊
	f.多罗霍伊Dorohoi.SetParent(f)
	
	f.赫尔勒乌Harlau = BHarlau赫尔勒乌
	f.赫尔勒乌Harlau.SetParent(f)
	
	f.霍京Hotin = BHotin霍京
	f.霍京Hotin.SetParent(f)
	
	f.雅西Iasi = BIasi雅西
	f.雅西Iasi.SetParent(f)
	
	f.温盖尼Ungheni = BUngheni温盖尼
	f.温盖尼Ungheni.SetParent(f)
	
	f.沃尔泽雷什蒂Varzaresti = BVarzaresti沃尔泽雷什蒂
	f.沃尔泽雷什蒂Varzaresti.SetParent(f)
	
}
