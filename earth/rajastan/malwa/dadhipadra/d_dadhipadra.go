package dadhipadra

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/dadhipadra/dadhipadra"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/dadhipadra/dasapura"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/dadhipadra/dhara"
	"github.com/thalesfu/CK2Commands/earth/rajastan/malwa/dadhipadra/ujjayini"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DadhipadraDuke interface {
    feud.Duke
    CDadhipadra陀提波陀罗() 	dadhipadra.DadhipadraCounty
    CDasapura陀舍补罗() 	dasapura.DasapuraCounty
    CDhara陀罗() 	dhara.DharaCounty
    CUjjayini邬阇衍那() 	ujjayini.UjjayiniCounty
}

type 阿般提DadhipadraDuke struct {
	feud.BaseDuke
	陀提波陀罗Dadhipadra 	dadhipadra.DadhipadraCounty
	陀舍补罗Dasapura 	dasapura.DasapuraCounty
	陀罗Dhara 	dhara.DharaCounty
	邬阇衍那Ujjayini 	ujjayini.UjjayiniCounty
}

func (d *阿般提DadhipadraDuke) CDadhipadra陀提波陀罗() dadhipadra.DadhipadraCounty {
	return d.陀提波陀罗Dadhipadra
}
    
func (d *阿般提DadhipadraDuke) CDasapura陀舍补罗() dasapura.DasapuraCounty {
	return d.陀舍补罗Dasapura
}
    
func (d *阿般提DadhipadraDuke) CDhara陀罗() dhara.DharaCounty {
	return d.陀罗Dhara
}
    
func (d *阿般提DadhipadraDuke) CUjjayini邬阇衍那() ujjayini.UjjayiniCounty {
	return d.邬阇衍那Ujjayini
}
    
var DDadhipadra阿般提 DadhipadraDuke = &阿般提DadhipadraDuke{}

func init() {
	f := DDadhipadra阿般提.(*阿般提DadhipadraDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "dadhipadra",
		TitleName: "阿般提",
		TitleCode: "d_dadhipadra",
		Counties:  map[string]feud.County{},
	}

	f.陀提波陀罗Dadhipadra = dadhipadra.CDadhipadra陀提波陀罗
	f.陀提波陀罗Dadhipadra.SetParent(f)
	
	f.陀舍补罗Dasapura = dasapura.CDasapura陀舍补罗
	f.陀舍补罗Dasapura.SetParent(f)
	
	f.陀罗Dhara = dhara.CDhara陀罗
	f.陀罗Dhara.SetParent(f)
	
	f.邬阇衍那Ujjayini = ujjayini.CUjjayini邬阇衍那
	f.邬阇衍那Ujjayini.SetParent(f)
	
}
