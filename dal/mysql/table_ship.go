package mysql

import (
	"log"
)

type ship struct {
	MMSI int64
}

func GetShipID() []int64 {
	rows, err := DB.Query("SELECT * from ship")
	if err != nil {
		log.Println("查询出错了")
		return []int64{}
	}
	shipIDs := make([]int64, 0)
	for rows.Next() {
		var info ship
		err := rows.Scan(&info.MMSI)
		if err != nil {
			log.Println("rows fail")
			break
		}
		shipIDs = append(shipIDs, info.MMSI)
	}
	return shipIDs
}
