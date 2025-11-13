package pegawai

type Pegawai struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
	Jabatan string `json:"jabatan"`
}

var PegawaiDB []Pegawai
var NextID = 1
