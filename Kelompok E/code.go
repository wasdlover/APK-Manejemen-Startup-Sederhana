package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const maxData = 100

type AnggotaTim struct {
	Nama  string
	Peran string
}

type Oger struct {
	Nama           string
	Bidang         string
	TahunBerdiri   int
	TotalPendanaan int
	Tim            [10]AnggotaTim
	JumlahAnggota  int
}

var daftarStartup [maxData]Oger
var jumlahStartup int

var reader = bufio.NewReader(os.Stdin)

// Util Input
func inputString(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func inputInt(prompt string) int {
	for {
		str := inputString(prompt)
		num, err := strconv.Atoi(str)
		if err == nil {
			return num
		}
		fmt.Println("Input harus berupa angka.")
	}
}

// Tambah Startup
func tambahStartup(data *[maxData]Oger, jumlah *int, s Oger) {
	if *jumlah < maxData {
		data[*jumlah] = s
		*jumlah++
		fmt.Println("Startup berhasil ditambahkan.")
	} else {
		fmt.Println("Kapasitas data penuh.")
	}
}

// Cari Startup
func cariStartupByNama(data [maxData]Oger, jumlah int, nama string) int {
	for i := 0; i < jumlah; i++ {
		if strings.EqualFold(data[i].Nama, nama) {
			return i
		}
	}
	return -1
}

// Ubah Startup
func ubahStartup(data *[maxData]Oger, idx int, s Oger) {
	data[idx] = s
	fmt.Println("Data berhasil diubah.")
}

// Hapus Startup
func hapusStartup(data *[maxData]Oger, jumlah *int, idx int) {
	for i := idx; i < *jumlah-1; i++ {
		data[i] = data[i+1]
	}
	*jumlah--
	fmt.Println("Data berhasil dihapus.")
}

// Tambah Anggota
func tambahAnggota(data *[maxData]Oger, idx int, anggota AnggotaTim) {
	if data[idx].JumlahAnggota < len(data[idx].Tim) {
		data[idx].Tim[data[idx].JumlahAnggota] = anggota
		data[idx].JumlahAnggota++
		fmt.Println("Anggota berhasil ditambahkan.")
	} else {
		fmt.Println("Tim sudah penuh.")
	}
}

// Selection Sort
func selectionSortPendanaan(data *[maxData]Oger, jumlah int, ascending bool) {
	for i := 0; i < jumlah-1; i++ {
		idx := i
		for j := i + 1; j < jumlah; j++ {
			if (ascending && data[j].TotalPendanaan < data[idx].TotalPendanaan) ||
				(!ascending && data[j].TotalPendanaan > data[idx].TotalPendanaan) {
				idx = j
			}
		}
		data[i], data[idx] = data[idx], data[i]
	}
}

// Insertion Sort
func insertionSortTahun(data *[maxData]Oger, jumlah int, ascending bool) {
	for i := 1; i < jumlah; i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && ((ascending && data[j].TahunBerdiri > temp.TahunBerdiri) ||
			(!ascending && data[j].TahunBerdiri < temp.TahunBerdiri)) {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

// Laporan
func laporanPerKategori(data [maxData]Oger, jumlah int) {
	kategori := map[string]int{}
	for i := 0; i < jumlah; i++ {
		kategori[data[i].Bidang]++
	}
	fmt.Println("\n--- Laporan Jumlah Startup per Kategori ---")
	for bidang, total := range kategori {
		fmt.Printf("%s: %d startup\n", bidang, total)
	}
}

// Tampilkan Semua Startup
func tampilkanStartup(data [maxData]Oger, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data startup.")
		return
	}
	fmt.Println("\n--- Daftar Startup ---")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("[%d] %s | %s | Tahun: %d | Pendanaan: %d | Tim: %d anggota\n",
			i+1, data[i].Nama, data[i].Bidang, data[i].TahunBerdiri, data[i].TotalPendanaan, data[i].JumlahAnggota)
	}
}

func main() {
	for {
		fmt.Println("\n=== Menu Manajemen Startup ===")
		fmt.Println("1. Tambah Startup")
		fmt.Println("2. Tampilkan Semua Startup")
		fmt.Println("3. Ubah Startup")
		fmt.Println("4. Hapus Startup")
		fmt.Println("5. Tambah Anggota Tim")
		fmt.Println("6. Cari Startup (Nama)")
		fmt.Println("7. Urutkan berdasarkan Pendanaan")
		fmt.Println("8. Urutkan berdasarkan Tahun Berdiri")
		fmt.Println("9. Laporan Startup per Kategori")
		fmt.Println("10. Lihat Detail Startup")
		fmt.Println("0. Keluar")

		pilihan := inputInt("Pilihan: ")

		switch pilihan {
		case 0:
			fmt.Println("Terima kasih!")
			return
		case 1:
			var s Oger
			s.Nama = inputString("Nama Startup: ")
			s.Bidang = inputString("Bidang Usaha: ")
			s.TahunBerdiri = inputInt("Tahun Berdiri: ")
			s.TotalPendanaan = inputInt("Total Pendanaan: ")
			tambahStartup(&daftarStartup, &jumlahStartup, s)
		case 2:
			tampilkanStartup(daftarStartup, jumlahStartup)
		case 3:
			nama := inputString("Masukkan nama startup yang ingin diubah: ")
			idx := cariStartupByNama(daftarStartup, jumlahStartup, nama)
			if idx != -1 {
				var s Oger
				s.Nama = inputString("Nama Baru: ")
				s.Bidang = inputString("Bidang Baru: ")
				s.TahunBerdiri = inputInt("Tahun Berdiri Baru: ")
				s.TotalPendanaan = inputInt("Total Pendanaan Baru: ")
				ubahStartup(&daftarStartup, idx, s)
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 4:
			nama := inputString("Masukkan nama startup yang ingin dihapus: ")
			idx := cariStartupByNama(daftarStartup, jumlahStartup, nama)
			if idx != -1 {
				hapusStartup(&daftarStartup, &jumlahStartup, idx)
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 5:
			nama := inputString("Masukkan nama startup: ")
			idx := cariStartupByNama(daftarStartup, jumlahStartup, nama)
			if idx != -1 {
				jumlah := inputInt("Berapa anggota tim yang ingin ditambahkan? ")
				for i := 0; i < jumlah; i++ {
					fmt.Printf("\nAnggota ke-%d:\n", i+1)
					namaAnggota := inputString("Nama Anggota Tim: ")
					peran := inputString("Peran Anggota: ")

					if daftarStartup[idx].JumlahAnggota >= len(daftarStartup[idx].Tim) {
						fmt.Println("Tim sudah mencapai kapasitas maksimum.")
						break
					}

					tambahAnggota(&daftarStartup, idx, AnggotaTim{namaAnggota, peran})
				}
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}

		case 6:
			nama := inputString("Nama Startup yang dicari: ")
			idx := cariStartupByNama(daftarStartup, jumlahStartup, nama)
			if idx != -1 {
				fmt.Printf("Ditemukan: %s (%s)\n", daftarStartup[idx].Nama, daftarStartup[idx].Bidang)
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}
		case 7:
			asc := inputInt("Urutkan secara Ascending (1) atau Descending (0): ") == 1
			selectionSortPendanaan(&daftarStartup, jumlahStartup, asc)
			tampilkanStartup(daftarStartup, jumlahStartup)
		case 8:
			asc := inputInt("Urutkan secara Ascending (1) atau Descending (0): ") == 1
			insertionSortTahun(&daftarStartup, jumlahStartup, asc)
			tampilkanStartup(daftarStartup, jumlahStartup)
		case 9:
			laporanPerKategori(daftarStartup, jumlahStartup)
		case 10:
			nama := inputString("Masukkan nama startup yang ingin dilihat: ")
			idx := cariStartupByNama(daftarStartup, jumlahStartup, nama)
			if idx != -1 {
				s := daftarStartup[idx]
				fmt.Println("\n--- Detail Startup ---")
				fmt.Println("Nama:           ", s.Nama)
				fmt.Println("Bidang Usaha:   ", s.Bidang)
				fmt.Println("Tahun Berdiri:  ", s.TahunBerdiri)
				fmt.Println("Pendanaan:      ", s.TotalPendanaan)
				fmt.Println("Jumlah Tim:     ", s.JumlahAnggota)

				if s.JumlahAnggota > 0 {
					fmt.Println("\n-- Daftar Tim --")
					for i := 0; i < s.JumlahAnggota; i++ {
						fmt.Printf("[%d] %s - %s\n", i+1, s.Tim[i].Nama, s.Tim[i].Peran)
					}
				} else {
					fmt.Println("Belum ada anggota tim.")
				}
			} else {
				fmt.Println("Startup tidak ditemukan.")
			}

		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
