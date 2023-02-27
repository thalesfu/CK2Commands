package susa

import (
	"github.com/thalesfu/CK2Commands/earth/italy/italy/susa/ivrea"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/susa/monferrato"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/susa/piemonte"
	"github.com/thalesfu/CK2Commands/earth/italy/italy/susa/saluzzo"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SusaDuke interface {
    feud.Duke
    CIvrea伊夫雷亚() 	ivrea.IvreaCounty
    CMonferrato蒙费拉托() 	monferrato.MonferratoCounty
    CPiemonte皮埃蒙特() 	piemonte.PiemonteCounty
    CSaluzzo萨卢佐() 	saluzzo.SaluzzoCounty
}

type 苏萨SusaDuke struct {
	feud.BaseDuke
	伊夫雷亚Ivrea 	ivrea.IvreaCounty
	蒙费拉托Monferrato 	monferrato.MonferratoCounty
	皮埃蒙特Piemonte 	piemonte.PiemonteCounty
	萨卢佐Saluzzo 	saluzzo.SaluzzoCounty
}

func (d *苏萨SusaDuke) CIvrea伊夫雷亚() ivrea.IvreaCounty {
	return d.伊夫雷亚Ivrea
}
    
func (d *苏萨SusaDuke) CMonferrato蒙费拉托() monferrato.MonferratoCounty {
	return d.蒙费拉托Monferrato
}
    
func (d *苏萨SusaDuke) CPiemonte皮埃蒙特() piemonte.PiemonteCounty {
	return d.皮埃蒙特Piemonte
}
    
func (d *苏萨SusaDuke) CSaluzzo萨卢佐() saluzzo.SaluzzoCounty {
	return d.萨卢佐Saluzzo
}
    
var DSusa苏萨 SusaDuke = &苏萨SusaDuke{}

func init() {
	f := DSusa苏萨.(*苏萨SusaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "susa",
		TitleName: "苏萨",
		TitleCode: "d_susa",
		Counties:  map[string]feud.County{},
	}

	f.伊夫雷亚Ivrea = ivrea.CIvrea伊夫雷亚
	f.伊夫雷亚Ivrea.SetParent(f)
	
	f.蒙费拉托Monferrato = monferrato.CMonferrato蒙费拉托
	f.蒙费拉托Monferrato.SetParent(f)
	
	f.皮埃蒙特Piemonte = piemonte.CPiemonte皮埃蒙特
	f.皮埃蒙特Piemonte.SetParent(f)
	
	f.萨卢佐Saluzzo = saluzzo.CSaluzzo萨卢佐
	f.萨卢佐Saluzzo.SetParent(f)
	
}
