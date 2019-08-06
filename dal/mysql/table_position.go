package mysql

import (
	"log"
)

type positon struct {
	ID                    int64
	Message_Type          int64
	Repeat_Indicator      int64
	MMSI                  int64
	Navigation_Status     int64
	ROT                   int64
	SOG                   float64
	Position_Accuracy     int64
	Longitude             float64
	Latitude              float64
	COG                   float64
	HDG                   int64
	Time_stamp            int64
	Reserved_for_regional int64
	RAIM_flag             int64
	Year                  int64
	Month                 int64
	Day                   int64
	Hour                  int64
	Minute                int64
	Second                int64
}

func GetPosition(id string) []positon {
	rows, err := DB.Query("SELECT * from position where MMSI = ?", id)
	if err != nil {
		log.Println("查询出错了")
		return nil
	}
	positions := make([]positon, 0)
	for rows.Next() {
		var pos positon
		err := rows.Scan(
			&pos.ID, &pos.Message_Type, &pos.Repeat_Indicator, &pos.MMSI, &pos.Navigation_Status, &pos.ROT, &pos.SOG,
			&pos.Position_Accuracy, &pos.Longitude, &pos.Latitude, &pos.COG, &pos.HDG, &pos.Time_stamp, &pos.Reserved_for_regional,
			&pos.RAIM_flag, &pos.Year, &pos.Month, &pos.Day, &pos.Hour, &pos.Minute, &pos.Second)
		if err != nil {
			log.Println("rows fail")
			break
		}
		positions = append(positions, pos)
	}
	return positions
}
