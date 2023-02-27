package gauda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GaudaCounty interface {
    feud.County
    BAdinatha_mandir阿提那他神庙() 	feud.Barony
    BAgmahl阿迦摩诃() 	feud.Barony
    BKarnasubarna羯罗拏苏伐剌那() 	feud.Barony
    BKumhaul贡呼() 	feud.Barony
    BRaktamrittika络多末知() 	feud.Barony
    BTarapith多罗比他() 	feud.Barony
    BTeliagarhi谛梨阿迦罗() 	feud.Barony
}

type 乔荼GaudaCounty struct {
	feud.BaseCounty
	阿提那他神庙Adinatha_mandir 	feud.Barony
	阿迦摩诃Agmahl 	feud.Barony
	羯罗拏苏伐剌那Karnasubarna 	feud.Barony
	贡呼Kumhaul 	feud.Barony
	络多末知Raktamrittika 	feud.Barony
	多罗比他Tarapith 	feud.Barony
	谛梨阿迦罗Teliagarhi 	feud.Barony
}

func (c *乔荼GaudaCounty) BAdinatha_mandir阿提那他神庙() feud.Barony {
	return c.阿提那他神庙Adinatha_mandir
}
    
func (c *乔荼GaudaCounty) BAgmahl阿迦摩诃() feud.Barony {
	return c.阿迦摩诃Agmahl
}
    
func (c *乔荼GaudaCounty) BKarnasubarna羯罗拏苏伐剌那() feud.Barony {
	return c.羯罗拏苏伐剌那Karnasubarna
}
    
func (c *乔荼GaudaCounty) BKumhaul贡呼() feud.Barony {
	return c.贡呼Kumhaul
}
    
func (c *乔荼GaudaCounty) BRaktamrittika络多末知() feud.Barony {
	return c.络多末知Raktamrittika
}
    
func (c *乔荼GaudaCounty) BTarapith多罗比他() feud.Barony {
	return c.多罗比他Tarapith
}
    
func (c *乔荼GaudaCounty) BTeliagarhi谛梨阿迦罗() feud.Barony {
	return c.谛梨阿迦罗Teliagarhi
}
    
var CGauda乔荼 GaudaCounty = &乔荼GaudaCounty{}

func init() {
	f := CGauda乔荼.(*乔荼GaudaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1243",
		Title:     "gauda",
		TitleName: "乔荼",
		TitleCode: "c_gauda",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿提那他神庙Adinatha_mandir = BAdinatha_mandir阿提那他神庙
	f.阿提那他神庙Adinatha_mandir.SetParent(f)
	
	f.阿迦摩诃Agmahl = BAgmahl阿迦摩诃
	f.阿迦摩诃Agmahl.SetParent(f)
	
	f.羯罗拏苏伐剌那Karnasubarna = BKarnasubarna羯罗拏苏伐剌那
	f.羯罗拏苏伐剌那Karnasubarna.SetParent(f)
	
	f.贡呼Kumhaul = BKumhaul贡呼
	f.贡呼Kumhaul.SetParent(f)
	
	f.络多末知Raktamrittika = BRaktamrittika络多末知
	f.络多末知Raktamrittika.SetParent(f)
	
	f.多罗比他Tarapith = BTarapith多罗比他
	f.多罗比他Tarapith.SetParent(f)
	
	f.谛梨阿迦罗Teliagarhi = BTeliagarhi谛梨阿迦罗
	f.谛梨阿迦罗Teliagarhi.SetParent(f)
	
}
