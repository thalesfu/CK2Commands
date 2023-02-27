package jersika

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JersikaCounty interface {
    feud.County
    BAlene阿莱内() 	feud.Barony
    BAsote阿索特() 	feud.Barony
    BAutine奥季内() 	feud.Barony
    BJersika耶尔西卡() 	feud.Barony
    BKoknese科克内塞() 	feud.Barony
    BMarciena马尔齐埃纳() 	feud.Barony
    BNegeste内格斯特() 	feud.Barony
}

type 耶尔西卡JersikaCounty struct {
	feud.BaseCounty
	阿莱内Alene 	feud.Barony
	阿索特Asote 	feud.Barony
	奥季内Autine 	feud.Barony
	耶尔西卡Jersika 	feud.Barony
	科克内塞Koknese 	feud.Barony
	马尔齐埃纳Marciena 	feud.Barony
	内格斯特Negeste 	feud.Barony
}

func (c *耶尔西卡JersikaCounty) BAlene阿莱内() feud.Barony {
	return c.阿莱内Alene
}
    
func (c *耶尔西卡JersikaCounty) BAsote阿索特() feud.Barony {
	return c.阿索特Asote
}
    
func (c *耶尔西卡JersikaCounty) BAutine奥季内() feud.Barony {
	return c.奥季内Autine
}
    
func (c *耶尔西卡JersikaCounty) BJersika耶尔西卡() feud.Barony {
	return c.耶尔西卡Jersika
}
    
func (c *耶尔西卡JersikaCounty) BKoknese科克内塞() feud.Barony {
	return c.科克内塞Koknese
}
    
func (c *耶尔西卡JersikaCounty) BMarciena马尔齐埃纳() feud.Barony {
	return c.马尔齐埃纳Marciena
}
    
func (c *耶尔西卡JersikaCounty) BNegeste内格斯特() feud.Barony {
	return c.内格斯特Negeste
}
    
var CJersika耶尔西卡 JersikaCounty = &耶尔西卡JersikaCounty{}

func init() {
	f := CJersika耶尔西卡.(*耶尔西卡JersikaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1594",
		Title:     "jersika",
		TitleName: "耶尔西卡",
		TitleCode: "c_jersika",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿莱内Alene = BAlene阿莱内
	f.阿莱内Alene.SetParent(f)
	
	f.阿索特Asote = BAsote阿索特
	f.阿索特Asote.SetParent(f)
	
	f.奥季内Autine = BAutine奥季内
	f.奥季内Autine.SetParent(f)
	
	f.耶尔西卡Jersika = BJersika耶尔西卡
	f.耶尔西卡Jersika.SetParent(f)
	
	f.科克内塞Koknese = BKoknese科克内塞
	f.科克内塞Koknese.SetParent(f)
	
	f.马尔齐埃纳Marciena = BMarciena马尔齐埃纳
	f.马尔齐埃纳Marciena.SetParent(f)
	
	f.内格斯特Negeste = BNegeste内格斯特
	f.内格斯特Negeste.SetParent(f)
	
}
