package chikoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChikoiCounty interface {
    feud.County
    BBursomon布尔索蒙() 	feud.Barony
    BCheremkhovo切列姆霍沃() 	feud.Barony
    BChikoi赤苦() 	feud.Barony
    BChudotvortsa丘多特沃尔察() 	feud.Barony
    BShergol_dzhin舍尔戈利金() 	feud.Barony
    BShimbilik申比利克() 	feud.Barony
    BUrluk乌尔卢克() 	feud.Barony
}

type 赤苦ChikoiCounty struct {
	feud.BaseCounty
	布尔索蒙Bursomon 	feud.Barony
	切列姆霍沃Cheremkhovo 	feud.Barony
	赤苦Chikoi 	feud.Barony
	丘多特沃尔察Chudotvortsa 	feud.Barony
	舍尔戈利金Shergol_dzhin 	feud.Barony
	申比利克Shimbilik 	feud.Barony
	乌尔卢克Urluk 	feud.Barony
}

func (c *赤苦ChikoiCounty) BBursomon布尔索蒙() feud.Barony {
	return c.布尔索蒙Bursomon
}
    
func (c *赤苦ChikoiCounty) BCheremkhovo切列姆霍沃() feud.Barony {
	return c.切列姆霍沃Cheremkhovo
}
    
func (c *赤苦ChikoiCounty) BChikoi赤苦() feud.Barony {
	return c.赤苦Chikoi
}
    
func (c *赤苦ChikoiCounty) BChudotvortsa丘多特沃尔察() feud.Barony {
	return c.丘多特沃尔察Chudotvortsa
}
    
func (c *赤苦ChikoiCounty) BShergol_dzhin舍尔戈利金() feud.Barony {
	return c.舍尔戈利金Shergol_dzhin
}
    
func (c *赤苦ChikoiCounty) BShimbilik申比利克() feud.Barony {
	return c.申比利克Shimbilik
}
    
func (c *赤苦ChikoiCounty) BUrluk乌尔卢克() feud.Barony {
	return c.乌尔卢克Urluk
}
    
var CChikoi赤苦 ChikoiCounty = &赤苦ChikoiCounty{}

func init() {
	f := CChikoi赤苦.(*赤苦ChikoiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1924",
		Title:     "chikoi",
		TitleName: "赤苦",
		TitleCode: "c_chikoi",
		Baronies:  map[string]feud.Barony{},
	}

	f.布尔索蒙Bursomon = BBursomon布尔索蒙
	f.布尔索蒙Bursomon.SetParent(f)
	
	f.切列姆霍沃Cheremkhovo = BCheremkhovo切列姆霍沃
	f.切列姆霍沃Cheremkhovo.SetParent(f)
	
	f.赤苦Chikoi = BChikoi赤苦
	f.赤苦Chikoi.SetParent(f)
	
	f.丘多特沃尔察Chudotvortsa = BChudotvortsa丘多特沃尔察
	f.丘多特沃尔察Chudotvortsa.SetParent(f)
	
	f.舍尔戈利金Shergol_dzhin = BShergol_dzhin舍尔戈利金
	f.舍尔戈利金Shergol_dzhin.SetParent(f)
	
	f.申比利克Shimbilik = BShimbilik申比利克
	f.申比利克Shimbilik.SetParent(f)
	
	f.乌尔卢克Urluk = BUrluk乌尔卢克
	f.乌尔卢克Urluk.SetParent(f)
	
}
