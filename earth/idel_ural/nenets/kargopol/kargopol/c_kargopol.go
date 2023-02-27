package kargopol

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KargopolCounty interface {
    feud.County
    BIvanovo_kargopol伊万诺沃() 	feud.Barony
    BKargopol卡尔戈波尔() 	feud.Barony
    BKazakovo卡扎科沃() 	feud.Barony
    BLacha拉恰() 	feud.Barony
    BLyaga利亚加() 	feud.Barony
    BPesok佩索克() 	feud.Barony
    BVoloshka沃洛什卡() 	feud.Barony
}

type 卡尔戈波尔KargopolCounty struct {
	feud.BaseCounty
	伊万诺沃Ivanovo_kargopol 	feud.Barony
	卡尔戈波尔Kargopol 	feud.Barony
	卡扎科沃Kazakovo 	feud.Barony
	拉恰Lacha 	feud.Barony
	利亚加Lyaga 	feud.Barony
	佩索克Pesok 	feud.Barony
	沃洛什卡Voloshka 	feud.Barony
}

func (c *卡尔戈波尔KargopolCounty) BIvanovo_kargopol伊万诺沃() feud.Barony {
	return c.伊万诺沃Ivanovo_kargopol
}
    
func (c *卡尔戈波尔KargopolCounty) BKargopol卡尔戈波尔() feud.Barony {
	return c.卡尔戈波尔Kargopol
}
    
func (c *卡尔戈波尔KargopolCounty) BKazakovo卡扎科沃() feud.Barony {
	return c.卡扎科沃Kazakovo
}
    
func (c *卡尔戈波尔KargopolCounty) BLacha拉恰() feud.Barony {
	return c.拉恰Lacha
}
    
func (c *卡尔戈波尔KargopolCounty) BLyaga利亚加() feud.Barony {
	return c.利亚加Lyaga
}
    
func (c *卡尔戈波尔KargopolCounty) BPesok佩索克() feud.Barony {
	return c.佩索克Pesok
}
    
func (c *卡尔戈波尔KargopolCounty) BVoloshka沃洛什卡() feud.Barony {
	return c.沃洛什卡Voloshka
}
    
var CKargopol卡尔戈波尔 KargopolCounty = &卡尔戈波尔KargopolCounty{}

func init() {
	f := CKargopol卡尔戈波尔.(*卡尔戈波尔KargopolCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1819",
		Title:     "kargopol",
		TitleName: "卡尔戈波尔",
		TitleCode: "c_kargopol",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊万诺沃Ivanovo_kargopol = BIvanovo_kargopol伊万诺沃
	f.伊万诺沃Ivanovo_kargopol.SetParent(f)
	
	f.卡尔戈波尔Kargopol = BKargopol卡尔戈波尔
	f.卡尔戈波尔Kargopol.SetParent(f)
	
	f.卡扎科沃Kazakovo = BKazakovo卡扎科沃
	f.卡扎科沃Kazakovo.SetParent(f)
	
	f.拉恰Lacha = BLacha拉恰
	f.拉恰Lacha.SetParent(f)
	
	f.利亚加Lyaga = BLyaga利亚加
	f.利亚加Lyaga.SetParent(f)
	
	f.佩索克Pesok = BPesok佩索克
	f.佩索克Pesok.SetParent(f)
	
	f.沃洛什卡Voloshka = BVoloshka沃洛什卡
	f.沃洛什卡Voloshka.SetParent(f)
	
}
