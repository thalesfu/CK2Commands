package castilla

import (
	"github.com/thalesfu/CK2Commands/earth/spain/castille/castilla/asturias_de_santillana"
	"github.com/thalesfu/CK2Commands/earth/spain/castille/castilla/burgos"
	"github.com/thalesfu/CK2Commands/earth/spain/castille/castilla/osma"
	"github.com/thalesfu/CK2Commands/earth/spain/castille/castilla/soria"
	"github.com/thalesfu/CK2Commands/earth/spain/castille/castilla/valladolid"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CastillaDuke interface {
    feud.Duke
    CAsturias_de_santillana阿斯图里亚斯德桑蒂利亚纳() 	asturias_de_santillana.Asturias_de_santillanaCounty
    CBurgos布尔戈斯() 	burgos.BurgosCounty
    COsma奥斯马() 	osma.OsmaCounty
    CSoria索里亚() 	soria.SoriaCounty
    CValladolid巴利亚多利德() 	valladolid.ValladolidCounty
}

type 卡斯蒂利亚CastillaDuke struct {
	feud.BaseDuke
	阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillana 	asturias_de_santillana.Asturias_de_santillanaCounty
	布尔戈斯Burgos 	burgos.BurgosCounty
	奥斯马Osma 	osma.OsmaCounty
	索里亚Soria 	soria.SoriaCounty
	巴利亚多利德Valladolid 	valladolid.ValladolidCounty
}

func (d *卡斯蒂利亚CastillaDuke) CAsturias_de_santillana阿斯图里亚斯德桑蒂利亚纳() asturias_de_santillana.Asturias_de_santillanaCounty {
	return d.阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillana
}
    
func (d *卡斯蒂利亚CastillaDuke) CBurgos布尔戈斯() burgos.BurgosCounty {
	return d.布尔戈斯Burgos
}
    
func (d *卡斯蒂利亚CastillaDuke) COsma奥斯马() osma.OsmaCounty {
	return d.奥斯马Osma
}
    
func (d *卡斯蒂利亚CastillaDuke) CSoria索里亚() soria.SoriaCounty {
	return d.索里亚Soria
}
    
func (d *卡斯蒂利亚CastillaDuke) CValladolid巴利亚多利德() valladolid.ValladolidCounty {
	return d.巴利亚多利德Valladolid
}
    
var DCastilla卡斯蒂利亚 CastillaDuke = &卡斯蒂利亚CastillaDuke{}

func init() {
	f := DCastilla卡斯蒂利亚.(*卡斯蒂利亚CastillaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "castilla",
		TitleName: "卡斯蒂利亚",
		TitleCode: "d_castilla",
		Counties:  map[string]feud.County{},
	}

	f.阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillana = asturias_de_santillana.CAsturias_de_santillana阿斯图里亚斯德桑蒂利亚纳
	f.阿斯图里亚斯德桑蒂利亚纳Asturias_de_santillana.SetParent(f)
	
	f.布尔戈斯Burgos = burgos.CBurgos布尔戈斯
	f.布尔戈斯Burgos.SetParent(f)
	
	f.奥斯马Osma = osma.COsma奥斯马
	f.奥斯马Osma.SetParent(f)
	
	f.索里亚Soria = soria.CSoria索里亚
	f.索里亚Soria.SetParent(f)
	
	f.巴利亚多利德Valladolid = valladolid.CValladolid巴利亚多利德
	f.巴利亚多利德Valladolid.SetParent(f)
	
}
