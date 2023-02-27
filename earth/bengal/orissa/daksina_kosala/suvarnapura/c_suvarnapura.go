package suvarnapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SuvarnapuraCounty interface {
    feud.County
    BAgalpur阏伽罗补罗() 	feud.Barony
    BBargarh跋罗姞利呬() 	feud.Barony
    BKalahandi迦罗汉持() 	feud.Barony
    BPatna波吒那() 	feud.Barony
    BSuvarnapura苏伐剌那补罗() 	feud.Barony
    BVinitapura毗尼多补罗() 	feud.Barony
    BYajatinagara耶阇底那揭罗() 	feud.Barony
}

type 苏伐剌那补罗SuvarnapuraCounty struct {
	feud.BaseCounty
	阏伽罗补罗Agalpur 	feud.Barony
	跋罗姞利呬Bargarh 	feud.Barony
	迦罗汉持Kalahandi 	feud.Barony
	波吒那Patna 	feud.Barony
	苏伐剌那补罗Suvarnapura 	feud.Barony
	毗尼多补罗Vinitapura 	feud.Barony
	耶阇底那揭罗Yajatinagara 	feud.Barony
}

func (c *苏伐剌那补罗SuvarnapuraCounty) BAgalpur阏伽罗补罗() feud.Barony {
	return c.阏伽罗补罗Agalpur
}
    
func (c *苏伐剌那补罗SuvarnapuraCounty) BBargarh跋罗姞利呬() feud.Barony {
	return c.跋罗姞利呬Bargarh
}
    
func (c *苏伐剌那补罗SuvarnapuraCounty) BKalahandi迦罗汉持() feud.Barony {
	return c.迦罗汉持Kalahandi
}
    
func (c *苏伐剌那补罗SuvarnapuraCounty) BPatna波吒那() feud.Barony {
	return c.波吒那Patna
}
    
func (c *苏伐剌那补罗SuvarnapuraCounty) BSuvarnapura苏伐剌那补罗() feud.Barony {
	return c.苏伐剌那补罗Suvarnapura
}
    
func (c *苏伐剌那补罗SuvarnapuraCounty) BVinitapura毗尼多补罗() feud.Barony {
	return c.毗尼多补罗Vinitapura
}
    
func (c *苏伐剌那补罗SuvarnapuraCounty) BYajatinagara耶阇底那揭罗() feud.Barony {
	return c.耶阇底那揭罗Yajatinagara
}
    
var CSuvarnapura苏伐剌那补罗 SuvarnapuraCounty = &苏伐剌那补罗SuvarnapuraCounty{}

func init() {
	f := CSuvarnapura苏伐剌那补罗.(*苏伐剌那补罗SuvarnapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1230",
		Title:     "suvarnapura",
		TitleName: "苏伐剌那补罗",
		TitleCode: "c_suvarnapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.阏伽罗补罗Agalpur = BAgalpur阏伽罗补罗
	f.阏伽罗补罗Agalpur.SetParent(f)
	
	f.跋罗姞利呬Bargarh = BBargarh跋罗姞利呬
	f.跋罗姞利呬Bargarh.SetParent(f)
	
	f.迦罗汉持Kalahandi = BKalahandi迦罗汉持
	f.迦罗汉持Kalahandi.SetParent(f)
	
	f.波吒那Patna = BPatna波吒那
	f.波吒那Patna.SetParent(f)
	
	f.苏伐剌那补罗Suvarnapura = BSuvarnapura苏伐剌那补罗
	f.苏伐剌那补罗Suvarnapura.SetParent(f)
	
	f.毗尼多补罗Vinitapura = BVinitapura毗尼多补罗
	f.毗尼多补罗Vinitapura.SetParent(f)
	
	f.耶阇底那揭罗Yajatinagara = BYajatinagara耶阇底那揭罗
	f.耶阇底那揭罗Yajatinagara.SetParent(f)
	
}
