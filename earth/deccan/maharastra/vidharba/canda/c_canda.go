package canda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CandaCounty interface {
    feud.County
    BBallarpur跋剌罗补罗() 	feud.Barony
    BCanda旃陀() 	feud.Barony
    BGadchiroli加德奇罗利() 	feud.Barony
    BKhunabat丘那跋() 	feud.Barony
    BMathila摩底陀() 	feud.Barony
    BNagbhid纳吉诃德() 	feud.Barony
    BRajura拉久拉() 	feud.Barony
}

type 旃陀CandaCounty struct {
	feud.BaseCounty
	跋剌罗补罗Ballarpur 	feud.Barony
	旃陀Canda 	feud.Barony
	加德奇罗利Gadchiroli 	feud.Barony
	丘那跋Khunabat 	feud.Barony
	摩底陀Mathila 	feud.Barony
	纳吉诃德Nagbhid 	feud.Barony
	拉久拉Rajura 	feud.Barony
}

func (c *旃陀CandaCounty) BBallarpur跋剌罗补罗() feud.Barony {
	return c.跋剌罗补罗Ballarpur
}
    
func (c *旃陀CandaCounty) BCanda旃陀() feud.Barony {
	return c.旃陀Canda
}
    
func (c *旃陀CandaCounty) BGadchiroli加德奇罗利() feud.Barony {
	return c.加德奇罗利Gadchiroli
}
    
func (c *旃陀CandaCounty) BKhunabat丘那跋() feud.Barony {
	return c.丘那跋Khunabat
}
    
func (c *旃陀CandaCounty) BMathila摩底陀() feud.Barony {
	return c.摩底陀Mathila
}
    
func (c *旃陀CandaCounty) BNagbhid纳吉诃德() feud.Barony {
	return c.纳吉诃德Nagbhid
}
    
func (c *旃陀CandaCounty) BRajura拉久拉() feud.Barony {
	return c.拉久拉Rajura
}
    
var CCanda旃陀 CandaCounty = &旃陀CandaCounty{}

func init() {
	f := CCanda旃陀.(*旃陀CandaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1269",
		Title:     "canda",
		TitleName: "旃陀",
		TitleCode: "c_canda",
		Baronies:  map[string]feud.Barony{},
	}

	f.跋剌罗补罗Ballarpur = BBallarpur跋剌罗补罗
	f.跋剌罗补罗Ballarpur.SetParent(f)
	
	f.旃陀Canda = BCanda旃陀
	f.旃陀Canda.SetParent(f)
	
	f.加德奇罗利Gadchiroli = BGadchiroli加德奇罗利
	f.加德奇罗利Gadchiroli.SetParent(f)
	
	f.丘那跋Khunabat = BKhunabat丘那跋
	f.丘那跋Khunabat.SetParent(f)
	
	f.摩底陀Mathila = BMathila摩底陀
	f.摩底陀Mathila.SetParent(f)
	
	f.纳吉诃德Nagbhid = BNagbhid纳吉诃德
	f.纳吉诃德Nagbhid.SetParent(f)
	
	f.拉久拉Rajura = BRajura拉久拉
	f.拉久拉Rajura.SetParent(f)
	
}
