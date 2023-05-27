package web

type ResponseStok struct {
	Bahan string
	Total int
}

type RequestStok struct {
	Bahan string
	Total int
}

type ResponseDetailStok struct {
	Bahan string
	Type  string
	Total int
}
