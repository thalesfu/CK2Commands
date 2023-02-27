package rohana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RohanaCounty interface {
    feud.County
    BBuduruvagala博杜瓦迦拉() 	feud.Barony
    BDevanagara提婆那揭罗() 	feud.Barony
    BDighavapi迭伽婆比() 	feud.Barony
    BGampala坎帕拉() 	feud.Barony
    BKajaragama迦阇罗伽摩() 	feud.Barony
    BMahagama摩诃伽摩() 	feud.Barony
    BMahanagakula摩诃那伽俱罗() 	feud.Barony
    BRohana卢诃拿() 	feud.Barony
}

type 卢诃拿RohanaCounty struct {
	feud.BaseCounty
	博杜瓦迦拉Buduruvagala 	feud.Barony
	提婆那揭罗Devanagara 	feud.Barony
	迭伽婆比Dighavapi 	feud.Barony
	坎帕拉Gampala 	feud.Barony
	迦阇罗伽摩Kajaragama 	feud.Barony
	摩诃伽摩Mahagama 	feud.Barony
	摩诃那伽俱罗Mahanagakula 	feud.Barony
	卢诃拿Rohana 	feud.Barony
}

func (c *卢诃拿RohanaCounty) BBuduruvagala博杜瓦迦拉() feud.Barony {
	return c.博杜瓦迦拉Buduruvagala
}
    
func (c *卢诃拿RohanaCounty) BDevanagara提婆那揭罗() feud.Barony {
	return c.提婆那揭罗Devanagara
}
    
func (c *卢诃拿RohanaCounty) BDighavapi迭伽婆比() feud.Barony {
	return c.迭伽婆比Dighavapi
}
    
func (c *卢诃拿RohanaCounty) BGampala坎帕拉() feud.Barony {
	return c.坎帕拉Gampala
}
    
func (c *卢诃拿RohanaCounty) BKajaragama迦阇罗伽摩() feud.Barony {
	return c.迦阇罗伽摩Kajaragama
}
    
func (c *卢诃拿RohanaCounty) BMahagama摩诃伽摩() feud.Barony {
	return c.摩诃伽摩Mahagama
}
    
func (c *卢诃拿RohanaCounty) BMahanagakula摩诃那伽俱罗() feud.Barony {
	return c.摩诃那伽俱罗Mahanagakula
}
    
func (c *卢诃拿RohanaCounty) BRohana卢诃拿() feud.Barony {
	return c.卢诃拿Rohana
}
    
var CRohana卢诃拿 RohanaCounty = &卢诃拿RohanaCounty{}

func init() {
	f := CRohana卢诃拿.(*卢诃拿RohanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1279",
		Title:     "rohana",
		TitleName: "卢诃拿",
		TitleCode: "c_rohana",
		Baronies:  map[string]feud.Barony{},
	}

	f.博杜瓦迦拉Buduruvagala = BBuduruvagala博杜瓦迦拉
	f.博杜瓦迦拉Buduruvagala.SetParent(f)
	
	f.提婆那揭罗Devanagara = BDevanagara提婆那揭罗
	f.提婆那揭罗Devanagara.SetParent(f)
	
	f.迭伽婆比Dighavapi = BDighavapi迭伽婆比
	f.迭伽婆比Dighavapi.SetParent(f)
	
	f.坎帕拉Gampala = BGampala坎帕拉
	f.坎帕拉Gampala.SetParent(f)
	
	f.迦阇罗伽摩Kajaragama = BKajaragama迦阇罗伽摩
	f.迦阇罗伽摩Kajaragama.SetParent(f)
	
	f.摩诃伽摩Mahagama = BMahagama摩诃伽摩
	f.摩诃伽摩Mahagama.SetParent(f)
	
	f.摩诃那伽俱罗Mahanagakula = BMahanagakula摩诃那伽俱罗
	f.摩诃那伽俱罗Mahanagakula.SetParent(f)
	
	f.卢诃拿Rohana = BRohana卢诃拿
	f.卢诃拿Rohana.SetParent(f)
	
}
