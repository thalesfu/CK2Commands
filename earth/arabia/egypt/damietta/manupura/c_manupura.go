package manupura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ManupuraCounty interface {
    feud.County
    BAthribis阿提利比斯() 	feud.Barony
    BBusiris布西里斯() 	feud.Barony
    BMalhalla迈哈莱() 	feud.Barony
    BMansoura曼苏雷() 	feud.Barony
    BSais塞易斯() 	feud.Barony
    BSamannud萨曼努德() 	feud.Barony
    BTanta坦泰() 	feud.Barony
}

type 迈哈莱ManupuraCounty struct {
	feud.BaseCounty
	阿提利比斯Athribis 	feud.Barony
	布西里斯Busiris 	feud.Barony
	迈哈莱Malhalla 	feud.Barony
	曼苏雷Mansoura 	feud.Barony
	塞易斯Sais 	feud.Barony
	萨曼努德Samannud 	feud.Barony
	坦泰Tanta 	feud.Barony
}

func (c *迈哈莱ManupuraCounty) BAthribis阿提利比斯() feud.Barony {
	return c.阿提利比斯Athribis
}
    
func (c *迈哈莱ManupuraCounty) BBusiris布西里斯() feud.Barony {
	return c.布西里斯Busiris
}
    
func (c *迈哈莱ManupuraCounty) BMalhalla迈哈莱() feud.Barony {
	return c.迈哈莱Malhalla
}
    
func (c *迈哈莱ManupuraCounty) BMansoura曼苏雷() feud.Barony {
	return c.曼苏雷Mansoura
}
    
func (c *迈哈莱ManupuraCounty) BSais塞易斯() feud.Barony {
	return c.塞易斯Sais
}
    
func (c *迈哈莱ManupuraCounty) BSamannud萨曼努德() feud.Barony {
	return c.萨曼努德Samannud
}
    
func (c *迈哈莱ManupuraCounty) BTanta坦泰() feud.Barony {
	return c.坦泰Tanta
}
    
var CManupura迈哈莱 ManupuraCounty = &迈哈莱ManupuraCounty{}

func init() {
	f := CManupura迈哈莱.(*迈哈莱ManupuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "797",
		Title:     "manupura",
		TitleName: "迈哈莱",
		TitleCode: "c_manupura",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿提利比斯Athribis = BAthribis阿提利比斯
	f.阿提利比斯Athribis.SetParent(f)
	
	f.布西里斯Busiris = BBusiris布西里斯
	f.布西里斯Busiris.SetParent(f)
	
	f.迈哈莱Malhalla = BMalhalla迈哈莱
	f.迈哈莱Malhalla.SetParent(f)
	
	f.曼苏雷Mansoura = BMansoura曼苏雷
	f.曼苏雷Mansoura.SetParent(f)
	
	f.塞易斯Sais = BSais塞易斯
	f.塞易斯Sais.SetParent(f)
	
	f.萨曼努德Samannud = BSamannud萨曼努德
	f.萨曼努德Samannud.SetParent(f)
	
	f.坦泰Tanta = BTanta坦泰
	f.坦泰Tanta.SetParent(f)
	
}
