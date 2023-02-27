package figuig

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FiguigCounty interface {
    feud.County
    BBenhammouch本哈穆什() 	feud.Barony
    BBouarfa布阿尔费() 	feud.Barony
    BDrablia德拉布利亚() 	feud.Barony
    BFiguig菲吉格() 	feud.Barony
    BGanastel加纳斯泰尔() 	feud.Barony
    BIch伊什() 	feud.Barony
    BRandon朗东() 	feud.Barony
}

type 坦德拉拉FiguigCounty struct {
	feud.BaseCounty
	本哈穆什Benhammouch 	feud.Barony
	布阿尔费Bouarfa 	feud.Barony
	德拉布利亚Drablia 	feud.Barony
	菲吉格Figuig 	feud.Barony
	加纳斯泰尔Ganastel 	feud.Barony
	伊什Ich 	feud.Barony
	朗东Randon 	feud.Barony
}

func (c *坦德拉拉FiguigCounty) BBenhammouch本哈穆什() feud.Barony {
	return c.本哈穆什Benhammouch
}
    
func (c *坦德拉拉FiguigCounty) BBouarfa布阿尔费() feud.Barony {
	return c.布阿尔费Bouarfa
}
    
func (c *坦德拉拉FiguigCounty) BDrablia德拉布利亚() feud.Barony {
	return c.德拉布利亚Drablia
}
    
func (c *坦德拉拉FiguigCounty) BFiguig菲吉格() feud.Barony {
	return c.菲吉格Figuig
}
    
func (c *坦德拉拉FiguigCounty) BGanastel加纳斯泰尔() feud.Barony {
	return c.加纳斯泰尔Ganastel
}
    
func (c *坦德拉拉FiguigCounty) BIch伊什() feud.Barony {
	return c.伊什Ich
}
    
func (c *坦德拉拉FiguigCounty) BRandon朗东() feud.Barony {
	return c.朗东Randon
}
    
var CFiguig坦德拉拉 FiguigCounty = &坦德拉拉FiguigCounty{}

func init() {
	f := CFiguig坦德拉拉.(*坦德拉拉FiguigCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "837",
		Title:     "figuig",
		TitleName: "坦德拉拉",
		TitleCode: "c_figuig",
		Baronies:  map[string]feud.Barony{},
	}

	f.本哈穆什Benhammouch = BBenhammouch本哈穆什
	f.本哈穆什Benhammouch.SetParent(f)
	
	f.布阿尔费Bouarfa = BBouarfa布阿尔费
	f.布阿尔费Bouarfa.SetParent(f)
	
	f.德拉布利亚Drablia = BDrablia德拉布利亚
	f.德拉布利亚Drablia.SetParent(f)
	
	f.菲吉格Figuig = BFiguig菲吉格
	f.菲吉格Figuig.SetParent(f)
	
	f.加纳斯泰尔Ganastel = BGanastel加纳斯泰尔
	f.加纳斯泰尔Ganastel.SetParent(f)
	
	f.伊什Ich = BIch伊什
	f.伊什Ich.SetParent(f)
	
	f.朗东Randon = BRandon朗东
	f.朗东Randon.SetParent(f)
	
}
