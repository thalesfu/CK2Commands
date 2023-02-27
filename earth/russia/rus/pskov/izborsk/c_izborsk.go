package izborsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IzborskCounty interface {
    feud.County
    BIzborsk伊兹博尔斯克() 	feud.Barony
    BKamno卡姆诺() 	feud.Barony
    BKhokhly霍赫雷() 	feud.Barony
    BPechory佩乔雷() 	feud.Barony
    BPoleny波列内() 	feud.Barony
    BRodovoye罗多沃耶() 	feud.Barony
    BTupy图佩() 	feud.Barony
}

type 伊兹博尔斯克IzborskCounty struct {
	feud.BaseCounty
	伊兹博尔斯克Izborsk 	feud.Barony
	卡姆诺Kamno 	feud.Barony
	霍赫雷Khokhly 	feud.Barony
	佩乔雷Pechory 	feud.Barony
	波列内Poleny 	feud.Barony
	罗多沃耶Rodovoye 	feud.Barony
	图佩Tupy 	feud.Barony
}

func (c *伊兹博尔斯克IzborskCounty) BIzborsk伊兹博尔斯克() feud.Barony {
	return c.伊兹博尔斯克Izborsk
}
    
func (c *伊兹博尔斯克IzborskCounty) BKamno卡姆诺() feud.Barony {
	return c.卡姆诺Kamno
}
    
func (c *伊兹博尔斯克IzborskCounty) BKhokhly霍赫雷() feud.Barony {
	return c.霍赫雷Khokhly
}
    
func (c *伊兹博尔斯克IzborskCounty) BPechory佩乔雷() feud.Barony {
	return c.佩乔雷Pechory
}
    
func (c *伊兹博尔斯克IzborskCounty) BPoleny波列内() feud.Barony {
	return c.波列内Poleny
}
    
func (c *伊兹博尔斯克IzborskCounty) BRodovoye罗多沃耶() feud.Barony {
	return c.罗多沃耶Rodovoye
}
    
func (c *伊兹博尔斯克IzborskCounty) BTupy图佩() feud.Barony {
	return c.图佩Tupy
}
    
var CIzborsk伊兹博尔斯克 IzborskCounty = &伊兹博尔斯克IzborskCounty{}

func init() {
	f := CIzborsk伊兹博尔斯克.(*伊兹博尔斯克IzborskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1663",
		Title:     "izborsk",
		TitleName: "伊兹博尔斯克",
		TitleCode: "c_izborsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊兹博尔斯克Izborsk = BIzborsk伊兹博尔斯克
	f.伊兹博尔斯克Izborsk.SetParent(f)
	
	f.卡姆诺Kamno = BKamno卡姆诺
	f.卡姆诺Kamno.SetParent(f)
	
	f.霍赫雷Khokhly = BKhokhly霍赫雷
	f.霍赫雷Khokhly.SetParent(f)
	
	f.佩乔雷Pechory = BPechory佩乔雷
	f.佩乔雷Pechory.SetParent(f)
	
	f.波列内Poleny = BPoleny波列内
	f.波列内Poleny.SetParent(f)
	
	f.罗多沃耶Rodovoye = BRodovoye罗多沃耶
	f.罗多沃耶Rodovoye.SetParent(f)
	
	f.图佩Tupy = BTupy图佩
	f.图佩Tupy.SetParent(f)
	
}
