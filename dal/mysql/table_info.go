package mysql

import (
	"log"
)

type info struct {
	ID 	 					int64
	Navigation_Status 		int64
	MMSI 					int64
	AIS						int64
	IMO 					int64
	Call_Sign 				string
	Name 					string
	Ship_Type 				int64
	A 						int64
	B 						int64
	C 						int64
	D 						int64
	Length 					int64
	Width 					int64
	Position_Type 			int64
	ETA_Month 				int64
	ETA_Day 				int64
	ETA_Hour 				int64
	ETA_Minute 				int64
	Draft 					float64
	Destination 			string
	Year 					int64
	Month 					int64
	Day 					int64
	Hour 					int64
	Minute 					int64
	Second 					int64
}

func GetShipInfo(id string) []info {
	rows, err := DB.Query("SELECT * from info where MMSI = ?", id)
	if err != nil{
		log.Println("查询出错了")
		return nil
	}
	infos := make([]info, 0)
	for rows.Next() {
		var inf info
		err := rows.Scan(
			&inf.ID, &inf.Navigation_Status, &inf.MMSI, &inf.AIS, &inf.IMO, &inf.Call_Sign, &inf.Name,
			&inf.Ship_Type, &inf.A, &inf.B, &inf.C, &inf.D, &inf.Length, &inf.Width,
			&inf.Position_Type, &inf.ETA_Month, &inf.ETA_Day, &inf.ETA_Hour, &inf.ETA_Minute, &inf.Draft, &inf.Destination,
			&inf.Year, &inf.Month, &inf.Day, &inf.Hour, &inf.Minute, &inf.Second)
		if err != nil {
			log.Println("rows fail")
			break
		}
		infos = append(infos, inf)
	}
	return infos
}