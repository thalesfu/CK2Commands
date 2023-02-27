package kafirkot

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KafirkotCounty interface {
    feud.County
    BBilot比诺特() 	feud.Barony
    BDera_ghazi_khan德拉加济汗() 	feud.Barony
    BKafirkot卡菲尔科特() 	feud.Barony
    BMalot马诺特() 	feud.Barony
    BManiot马尼诃特() 	feud.Barony
    BRupsi卢斯() 	feud.Barony
    BSagra娑伽罗() 	feud.Barony
}

type 卡菲尔科特KafirkotCounty struct {
	feud.BaseCounty
	比诺特Bilot 	feud.Barony
	德拉加济汗Dera_ghazi_khan 	feud.Barony
	卡菲尔科特Kafirkot 	feud.Barony
	马诺特Malot 	feud.Barony
	马尼诃特Maniot 	feud.Barony
	卢斯Rupsi 	feud.Barony
	娑伽罗Sagra 	feud.Barony
}

func (c *卡菲尔科特KafirkotCounty) BBilot比诺特() feud.Barony {
	return c.比诺特Bilot
}
    
func (c *卡菲尔科特KafirkotCounty) BDera_ghazi_khan德拉加济汗() feud.Barony {
	return c.德拉加济汗Dera_ghazi_khan
}
    
func (c *卡菲尔科特KafirkotCounty) BKafirkot卡菲尔科特() feud.Barony {
	return c.卡菲尔科特Kafirkot
}
    
func (c *卡菲尔科特KafirkotCounty) BMalot马诺特() feud.Barony {
	return c.马诺特Malot
}
    
func (c *卡菲尔科特KafirkotCounty) BManiot马尼诃特() feud.Barony {
	return c.马尼诃特Maniot
}
    
func (c *卡菲尔科特KafirkotCounty) BRupsi卢斯() feud.Barony {
	return c.卢斯Rupsi
}
    
func (c *卡菲尔科特KafirkotCounty) BSagra娑伽罗() feud.Barony {
	return c.娑伽罗Sagra
}
    
var CKafirkot卡菲尔科特 KafirkotCounty = &卡菲尔科特KafirkotCounty{}

func init() {
	f := CKafirkot卡菲尔科特.(*卡菲尔科特KafirkotCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1375",
		Title:     "kafirkot",
		TitleName: "卡菲尔科特",
		TitleCode: "c_kafirkot",
		Baronies:  map[string]feud.Barony{},
	}

	f.比诺特Bilot = BBilot比诺特
	f.比诺特Bilot.SetParent(f)
	
	f.德拉加济汗Dera_ghazi_khan = BDera_ghazi_khan德拉加济汗
	f.德拉加济汗Dera_ghazi_khan.SetParent(f)
	
	f.卡菲尔科特Kafirkot = BKafirkot卡菲尔科特
	f.卡菲尔科特Kafirkot.SetParent(f)
	
	f.马诺特Malot = BMalot马诺特
	f.马诺特Malot.SetParent(f)
	
	f.马尼诃特Maniot = BManiot马尼诃特
	f.马尼诃特Maniot.SetParent(f)
	
	f.卢斯Rupsi = BRupsi卢斯
	f.卢斯Rupsi.SetParent(f)
	
	f.娑伽罗Sagra = BSagra娑伽罗
	f.娑伽罗Sagra.SetParent(f)
	
}
