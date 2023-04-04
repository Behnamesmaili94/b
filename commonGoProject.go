package main

func main() {

	getData()

}

func getData() {

	type str struct {
		name  string
		id    int
		value int
	}

	var setData = []str{
		{"ali", 1, 19000},
		{"behzad", 2, 29000},
		{"mohsen", 3, 39000},
		{"moslem", 4, 49000},
	}

	for i := 0; i < len(setData); i++ {

		println("name : ", setData[i].name)
		println("value : ", setData[i].value)
		println("id : ", setData[i].id)
	}

}
