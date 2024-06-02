package main

import "fmt"

const NMAX int = 40

type MataKuliah struct {
	Nama, Grade           string
	SKS                   int
	UTS, UAS, Quiz, Total float64
}

type Mahasiswa struct {
	NIM, Nama   string
	MataKuliahs [NMAX]MataKuliah
	JumlahMK    int
}

type arrMhs [NMAX]Mahasiswa

func inputMahasiswa(A *arrMhs, N *int) {
	fmt.Print("Masukkan jumlah Mahasiswa: ")
	fmt.Scan(N)
	for i := 0; i < *N; i++ {
		fmt.Print("Masukkan NIM dan Nama Mahasiswa: ")
		fmt.Scan(&A[i].NIM, &A[i].Nama)
	}
}

func editMahasiswa(A *arrMhs, N int) {
	var nim, namaBaru string
	fmt.Print("Masukkan NIM yang di edit: ")
	fmt.Scan(&nim)

	for i := 0; i < N; i++ {
		if A[i].NIM == nim {
			fmt.Print("Masukkan Nama Baru: ")
			fmt.Scan(&namaBaru)
			A[i].Nama = namaBaru
			fmt.Println("Nama Mahasiswa Berhasil di update.")
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan.")
}

func hapusMahasiswa(A *arrMhs, N *int) {
	var nim string
	fmt.Print("Masukkan NIM yang akan di hapus: ")
	fmt.Scan(&nim)

	for i := 0; i < *N; i++ {
		if A[i].NIM == nim {
			for j := i; j < *N-1; j++ {
				A[j] = A[j+1]
			}
			A[*N-1] = Mahasiswa{}
			(*N)--
			fmt.Println("Mahasiswa berhasil di hapus.")
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan.")
}

func tambahMataKuliah(A *arrMhs, N int, nim, namaMK string, sks int, uts, uas, quiz float64) {
	for i := 0; i < N; i++ {
		if A[i].NIM == nim {
			total := uts*0.4 + uas*0.4 + quiz*0.2
			grade := hitungGrade(total)
			A[i].MataKuliahs[A[i].JumlahMK].Nama = namaMK
			A[i].MataKuliahs[A[i].JumlahMK].Grade = grade
			A[i].MataKuliahs[A[i].JumlahMK].SKS = sks
			A[i].MataKuliahs[A[i].JumlahMK].UTS = uts
			A[i].MataKuliahs[A[i].JumlahMK].UAS = uas
			A[i].MataKuliahs[A[i].JumlahMK].Quiz = quiz
			A[i].MataKuliahs[A[i].JumlahMK].Total = total
			A[i].JumlahMK++
			fmt.Println("Mata Kuliah Berhasil ditambahkan")
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan")
}

func editMataKuliah(A *arrMhs, N int, nim, namaMK string) {
	var input string
	var name string
	var sks int
	var uts, uas, quiz float64
	for i := 0; i < N; i++ {
		if A[i].NIM == nim {
			for j := 0; j < A[i].JumlahMK; j++ {
				if A[i].MataKuliahs[j].Nama == namaMK {
					fmt.Println("Apa yang perlu di ubah?(Nama, SKS, UTS, UAS, Quiz)")
					fmt.Scan(&input)
					if input == "Nama" {
						fmt.Print("Nama baru:")
						fmt.Scan(&name)
						A[i].MataKuliahs[j].Nama = name
					} else if input == "SKS" {
						fmt.Print("SKS baru:")
						fmt.Scan(&sks)
						A[i].MataKuliahs[j].SKS = sks
					} else if input == "UTS" {
						fmt.Print("UTS Baru:")
						fmt.Scan(&uts)
						A[i].MataKuliahs[j].UTS = uts
						A[i].MataKuliahs[j].Total = A[i].MataKuliahs[j].UTS * 0.4
					} else if input == "UAS" {
						fmt.Print("UAS Baru: ")
						fmt.Scan(&uas)
						A[i].MataKuliahs[j].UAS = uas
						A[i].MataKuliahs[j].Total = A[i].MataKuliahs[j].UTS * 0.4
					} else if input == "Quiz" {
						fmt.Print("Quiz Baru: ")
						fmt.Scan(&quiz)
						A[i].MataKuliahs[j].Quiz = quiz
						A[i].MataKuliahs[j].Total = A[i].MataKuliahs[j].Quiz * 0.2
					} else {
						fmt.Println("Input tidak valid.")
					}
				}
				A[i].MataKuliahs[j].Total = A[i].MataKuliahs[j].UTS*0.4 + A[i].MataKuliahs[j].UAS*0.4 + A[i].MataKuliahs[j].Quiz*0.2
				A[i].MataKuliahs[j].Grade = hitungGrade(A[i].MataKuliahs[j].Total)
				fmt.Println("Data Mata Kuliah berhasil diubah.")
				return
			}
		}
	}
	fmt.Println("Mahasiswa atau Mata Kuliah tidak ditemukan.")
}

func hapusMataKuliah(A *arrMhs, N int, nim, namaMK string) {
	for i := 0; i < N; i++ {
		if A[i].NIM == nim {
			for j := 0; j < A[i].JumlahMK; j++ {
				if A[i].MataKuliahs[j].Nama == namaMK {
					for k := j; k < A[i].JumlahMK-1; k++ {
						A[i].MataKuliahs[k] = A[i].MataKuliahs[k+1]
					}
					A[i].MataKuliahs[A[i].JumlahMK-1] = MataKuliah{}
					A[i].JumlahMK--
					fmt.Println("Mata Kuliah berhasil dihapus.")
					return
				}
			}
		}
	}
	fmt.Println("Mahasiswa atau Mata Kuliah tidak ditemukan.")
}

func hitungGrade(total float64) string {
	if total >= 80 {
		return "A"
	} else if total >= 70 {
		return "B"
	} else if total >= 65 {
		return "C"
	} else if total >= 50 {
		return "D"
	} else {
		return "E"
	}
}

func tampilkanDaftarMahasiswaBerdasarkanMataKuliah(A *arrMhs, N int, namaMK string) {
	var mahasiswaDitemukan bool
	fmt.Printf("Daftar Mahasiswa yang mengambil Mata Kuliah %s:\n", namaMK)
	for i := 0; i < N; i++ {
		for j := 0; j < A[i].JumlahMK; j++ {
			if A[i].MataKuliahs[j].Nama == namaMK {
				fmt.Printf("NIM : %s, Nama: %s\n", A[i].NIM, A[i].Nama)
				mahasiswaDitemukan = true
			}
		}
	}
	if !mahasiswaDitemukan {
		fmt.Println("Mahasiswa dengan Mata Kuliah", namaMK, "tidak ditemukan.")
	}
}

func tampilkanDaftarMataKuliahBerdasarkanMahasiswa(A *arrMhs, N int, nim string) {
	for i := 0; i < N; i++ {
		if A[i].NIM == nim {
			fmt.Printf("Daftar Mata Kuliah yang diambil oleh Mahasiswa %s (%s):\n", A[i].Nama, nim)
			for j := 0; j < A[i].JumlahMK; j++ {
				mk := A[i].MataKuliahs[j]
				fmt.Printf("Mata Kuliah: %s, Nilai: %.2f, SKS: %d\n", mk.Nama, mk.Total, mk.SKS)
			}
			return
		}
	}
	fmt.Println("Mahasiswa dengan NIM", nim, "tidak ditemukan.")
}

func tampilkanMahasiswaTerurutBerdarakanSKS(A *arrMhs, N int) {
	for i := 1; i < N; i++ {
		j := i
		for j > 0 && A[j-1].MataKuliahs[0].SKS < A[j].MataKuliahs[0].SKS {
			A[j], A[j-1] = A[j-1], A[j]
			j = j - 1
		}
	}

	fmt.Println("Mahasiswa terurut berdasarkan SKS yang diambil:")
	for i := 0; i < N; i++ {
		if A[i].JumlahMK > 0 {
			fmt.Printf("NIM: %s, Nama: %s, SKS: %d\n", A[i].NIM, A[i].Nama, A[i].MataKuliahs[0].SKS)
		}
	}
}

func tampilkanTranskripNilai(A *arrMhs, N int, nim string) {
	for i := 0; i < N; i++ {
		if A[i].NIM == nim {
			fmt.Printf("Transkrip Nilai Mahasiswa %s:\n", A[i].Nama)
			for j := 0; j < A[i].JumlahMK; j++ {
				mk := A[i].MataKuliahs[j]
				fmt.Printf("Mata Kuliah: %s, SKS: %d, UTS: %.2f, UAS: %.2f, Quiz: %.2f, Total: %.2f, Grade: %s\n", mk.Nama, mk.SKS, mk.UTS, mk.UAS, mk.Quiz, mk.Total, mk.Grade)
			}
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan.")
}

func urutBerdasarkanNilaiTertinggi(A *arrMhs, N int) {
	for i := 0; i < N-1; i++ {
		maxIdx := i
		for j := i + 1; j < N; j++ {
			if A[j].MataKuliahs[0].Total > A[maxIdx].MataKuliahs[0].Total {
				maxIdx = j
			}
		}
		A[i], A[maxIdx] = A[maxIdx], A[i]
	}

	fmt.Println("Daftar Mahasiswa urut berdasarkan nilai tertinggi:")
	for i := 0; i < N; i++ {
		if A[i].JumlahMK > 0 {
			fmt.Printf("NIM: %s, Nama: %s, Nilai: %.2f\n", A[i].NIM, A[i].Nama, A[i].MataKuliahs[0].Total)
		}
	}
}

func urutBerdasarkanNilaiTerendah(A *arrMhs, N int) {
	for i := 0; i < N-1; i++ {
		minIdx := i
		for j := i + 1; j < N; j++ {
			if A[j].MataKuliahs[0].Total < A[minIdx].MataKuliahs[0].Total {
				minIdx = j
			}
		}
		A[i], A[minIdx] = A[minIdx], A[i]
	}

	fmt.Println("Daftar Mahasiswa urut berdasarkan nilai terendah:")
	for i := 0; i < N; i++ {
		if A[i].JumlahMK > 0 {
			fmt.Printf("NIM: %s, Nama: %s, Nilai: %.2f\n", A[i].NIM, A[i].Nama, A[i].MataKuliahs[0].Total)
		}
	}
}

func main() {
	var A arrMhs
	var N int

	for {
		fmt.Println("\nKategori:")
		fmt.Println("1. Input Mahasiswa")
		fmt.Println("2. Edit Mahasiswa")
		fmt.Println("3. Hapus Mahasiswa")
		fmt.Println("4. Tambah Mata Kuliah")
		fmt.Println("5. Edit Mata Kuliah")
		fmt.Println("6. Hapus Mata Kuliah")
		fmt.Println("7. Tampilkan Daftar Mahasiswa Berdasarkan Mata Kuliah")
		fmt.Println("8. Tampilkan Daftar Mata Kuliah Berdasarkan Mahasiswa")
		fmt.Println("9. Tampilkan Mahasiswa Terurut Berdasarkan SKS")
		fmt.Println("10. Tampilkan Mahasiswa Terurut Berdasarkan Nilai Tertinggi")
		fmt.Println("11. Tampilkan Mahasiswa Terurut Berdasarkan Nilai Terendah")
		fmt.Println("12. Tampilkan Transkrip Nilai")
		fmt.Println("0. Exit")
		fmt.Print("Pilihan: ")

		var pilihan int
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			inputMahasiswa(&A, &N)
		} else if pilihan == 2 {
			editMahasiswa(&A, N)
		} else if pilihan == 3 {
			hapusMahasiswa(&A, &N)
		} else if pilihan == 4 {
			var nim, namaMK string
			var sks int
			var uts, uas, quiz float64
			fmt.Print("Masukkan NIM Mahasiswa: ")
			fmt.Scan(&nim)
			fmt.Print("Masukkan Nama Mata Kuliah: ")
			fmt.Scan(&namaMK)
			fmt.Print("Masukkan SKS: ")
			fmt.Scan(&sks)
			fmt.Print("Masukkan Nilai UTS: ")
			fmt.Scan(&uts)
			fmt.Print("Masukkan Nilai UAS: ")
			fmt.Scan(&uas)
			fmt.Print("Masukkan Nilai Quiz: ")
			fmt.Scan(&quiz)
			tambahMataKuliah(&A, N, nim, namaMK, sks, uts, uas, quiz)
		} else if pilihan == 5 {
			var nim, namaMK string
			fmt.Print("Masukkan NIM Mahasiswa: ")
			fmt.Scan(&nim)
			fmt.Print("Masukkan Nama Mata Kuliah: ")
			fmt.Scan(&namaMK)
			editMataKuliah(&A, N, nim, namaMK)
		} else if pilihan == 6 {
			var nim, namaMK string
			fmt.Print("Masukkan NIM Mahasiswa: ")
			fmt.Scan(&nim)
			fmt.Print("Masukkan Nama Mata Kuliah: ")
			fmt.Scan(&namaMK)
			hapusMataKuliah(&A, N, nim, namaMK)
		} else if pilihan == 7 {
			var namaMK string
			fmt.Print("Masukkan Nama Mata Kuliah: ")
			fmt.Scan(&namaMK)
			tampilkanDaftarMahasiswaBerdasarkanMataKuliah(&A, N, namaMK)
		} else if pilihan == 8 {
			var nim string
			fmt.Print("Masukkan NIM Mahasiswa: ")
			fmt.Scan(&nim)
			tampilkanDaftarMataKuliahBerdasarkanMahasiswa(&A, N, nim)
		} else if pilihan == 9 {
			tampilkanMahasiswaTerurutBerdarakanSKS(&A, N)
		} else if pilihan == 10 {
			urutBerdasarkanNilaiTertinggi(&A, N)
		} else if pilihan == 11 {
			urutBerdasarkanNilaiTerendah(&A, N)
		} else if pilihan == 12 {
			var nim string
			fmt.Print("Masukkan NIM Mahasiswa: ")
			fmt.Scan(&nim)
			tampilkanTranskripNilai(&A, N, nim)
		} else if pilihan == 0 {
			fmt.Println("Exiting...")
			return
		} else {
			fmt.Println("Pilihan tidak valid, coba lagi.")
		}
	}
}
