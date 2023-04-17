package db

func GetNewDbConnection() *map[int32]string {
	return &map[int32]string{1: "one", 2: "two", 3: "three"}
}
