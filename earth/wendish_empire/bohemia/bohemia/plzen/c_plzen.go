package plzen

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PlzenCounty interface {
    feud.County
    BDomazlice多马日利采() 	feud.Barony
    BKladruby克拉德鲁比() 	feud.Barony
    BKlatovy克拉托维() 	feud.Barony
    BPlasy普拉西() 	feud.Barony
    BPlzen普尔曾() 	feud.Barony
    BPrimda普日姆达() 	feud.Barony
    BStribo斯特日布罗() 	feud.Barony
}

type 普尔曾PlzenCounty struct {
	feud.BaseCounty
	多马日利采Domazlice 	feud.Barony
	克拉德鲁比Kladruby 	feud.Barony
	克拉托维Klatovy 	feud.Barony
	普拉西Plasy 	feud.Barony
	普尔曾Plzen 	feud.Barony
	普日姆达Primda 	feud.Barony
	斯特日布罗Stribo 	feud.Barony
}

func (c *普尔曾PlzenCounty) BDomazlice多马日利采() feud.Barony {
	return c.多马日利采Domazlice
}
    
func (c *普尔曾PlzenCounty) BKladruby克拉德鲁比() feud.Barony {
	return c.克拉德鲁比Kladruby
}
    
func (c *普尔曾PlzenCounty) BKlatovy克拉托维() feud.Barony {
	return c.克拉托维Klatovy
}
    
func (c *普尔曾PlzenCounty) BPlasy普拉西() feud.Barony {
	return c.普拉西Plasy
}
    
func (c *普尔曾PlzenCounty) BPlzen普尔曾() feud.Barony {
	return c.普尔曾Plzen
}
    
func (c *普尔曾PlzenCounty) BPrimda普日姆达() feud.Barony {
	return c.普日姆达Primda
}
    
func (c *普尔曾PlzenCounty) BStribo斯特日布罗() feud.Barony {
	return c.斯特日布罗Stribo
}
    
var CPlzen普尔曾 PlzenCounty = &普尔曾PlzenCounty{}

func init() {
	f := CPlzen普尔曾.(*普尔曾PlzenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "439",
		Title:     "plzen",
		TitleName: "普尔曾",
		TitleCode: "c_plzen",
		Baronies:  map[string]feud.Barony{},
	}

	f.多马日利采Domazlice = BDomazlice多马日利采
	f.多马日利采Domazlice.SetParent(f)
	
	f.克拉德鲁比Kladruby = BKladruby克拉德鲁比
	f.克拉德鲁比Kladruby.SetParent(f)
	
	f.克拉托维Klatovy = BKlatovy克拉托维
	f.克拉托维Klatovy.SetParent(f)
	
	f.普拉西Plasy = BPlasy普拉西
	f.普拉西Plasy.SetParent(f)
	
	f.普尔曾Plzen = BPlzen普尔曾
	f.普尔曾Plzen.SetParent(f)
	
	f.普日姆达Primda = BPrimda普日姆达
	f.普日姆达Primda.SetParent(f)
	
	f.斯特日布罗Stribo = BStribo斯特日布罗
	f.斯特日布罗Stribo.SetParent(f)
	
}
