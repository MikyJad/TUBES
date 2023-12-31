package main
import (
    "fmt"
)

func main() {
    var x, y, z, sk1, sk2, sk3, kodok, mau, mau2, mau3, lawan, roket, benar string
    var nyawa, nyawaku, ngurang int

    nyawa = 10
	nyawaku = 10
    
    fmt.Println("Selamat datang di gerbang Dunia Isekai")
    fmt.Println("Untuk melanjutkan perjalanan, kamu harus mengalahkan monster dulu")
    fmt.Println("Apa kamu siap?")

    fmt.Scan(&x)
    if x == "ya" {
        fmt.Println("")
        fmt.Println("")
        fmt.Println("HAHAHA")
        fmt.Println("")
        fmt.Println("Lawanmu adalah Kodok Acumalaka")
        fmt.Println("")
        fmt.Println("Siapkan petarung terbaikmu!")

        fmt.Println("Pilih Petarung : 1 / 2 / 3")
        fmt.Println("1. Kucin")
        fmt.Println("2. Sigit Rendang")
        fmt.Println("3. Soldier")
        fmt.Println("*Setiap petarung memiliki skill berbeda")
        fmt.Scan(&y)

        if y == "1" {
            sk1 = "Haheng"
            sk2 = "Huh"
            sk3 = "Meow"

            fmt.Println("")
            fmt.Println("")
            fmt.Println("Siap?")
            fmt.Println("")
            fmt.Println("Pertarungan dimulai HAHAHA")

            for nyawa > 0 && nyawaku > 0 {
                fmt.Println("")
                fmt.Println("")
                fmt.Println("Pilih Skill : 1 / 2 / 3")
                fmt.Println("1.", sk1)
                fmt.Println("2.", sk2)
                fmt.Println("3.", sk3)
                fmt.Scan(&z)
                fmt.Println("")

                if z == "1" {
                    ngurang = 3
                    fmt.Println("Kucin Use", sk1)
					fmt.Println("Efektif!!")
                } else if z == "2" {
                    ngurang = 1
                    fmt.Println("Kucin Use", sk2)
					fmt.Println("Malah bingung, pake yang lain!")
                } else if z == "3" {
                    ngurang = 1
                    fmt.Println("Kucin Use", sk3)
					fmt.Println("Ga Efektif, pake yang lain!")
                }

                nyawa = nyawa - ngurang
				nyawaku = nyawaku - 2
                fmt.Println("")
				
                if nyawa > 0 && nyawaku > 0 {
                    fmt.Println("Nyawa Kodok Acumalaka :", nyawa)
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Kodok Acumalaka use Backflip")
					fmt.Println("")
					fmt.Println("Nyawa Kucin :", nyawaku)
                } else if nyawa <= 0 && nyawaku > 0 || nyawa < nyawaku {
                    nyawa = 0
                    fmt.Println("Nyawa Kodok Acumalaka :", nyawa)

					
					fmt.Println("Kodok Acumalaka berhasil dikalahkan!")
					fmt.Println("")
					fmt.Println("Apakah kamu ingin menangkap Kodok Acumalaka?")
		
					fmt.Scan(&kodok)
					if kodok == "ya"{
						fmt.Println("")
						fmt.Println("Kamu memakai KepoBall")
						fmt.Println("")
						fmt.Println("Kodok Acumalaka berhasil ditangkap")
						fmt.Println("")
						fmt.Println("Kodok Acumalaka disimpan di tas")
					} else if kodok == "tidak" {
						fmt.Println("")
						fmt.Println("Okelah")
					}
		
					fmt.Println("")
					fmt.Println("Sekarang kamu bisa masuk ke Dunia Isekai")
					fmt.Println("")
					fmt.Println("Welcome!")
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Dibuat oleh Mikhael Restu Mahardhika")
					fmt.Println("103012300277")
					fmt.Println("Hehe TY ALL!")
                } else if nyawa > 0 && nyawaku <= 0 {
					fmt.Println("Kodok Acumalaka use Backflip")
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Kamu berhasil dikalahkan!")
					fmt.Println("Selamat tinggal!")
				} else if nyawa <= 0 && nyawaku <= 0 {
					fmt.Println("Kodok Acumalaka use Backflip")
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Kalian seri!!")
					fmt.Println("")
					fmt.Println("Kembalilah, bertarunglah lain waktu!")
				}


            }


        } else if y == "2" {
            sk1 = "Sniff"
            sk2 = "Leap"
            sk3 = "Jathilan"

            fmt.Println("")
            fmt.Println("Siap?")
            fmt.Println("")
            fmt.Println("Pertarungan dimulai HAHAHA")

            for nyawa > 0 && nyawaku > 0 {
                fmt.Println("")
                fmt.Println("")
                fmt.Println("Pilih Skill :")
                fmt.Println("1.", sk1)
                fmt.Println("2.", sk2)
                fmt.Println("3.", sk3)
                fmt.Scan(&z)
                fmt.Println("")

                if z == "1" {
                    ngurang = 1
                    fmt.Println("Sigit Rendang Use", sk1)
					fmt.Println("Ga Efektif, pake yang lain!")
                } else if z == "2" {
                    ngurang = 1
                    fmt.Println("Sigit Rendang Use", sk2)
					fmt.Println("Ga Efektif, pake yang lain!")
                } else if z == "3" {
                    ngurang = 3
                    fmt.Println("Sigit Rendang Use", sk3)
					fmt.Println("Efektif!!")
                }

                nyawa = nyawa - ngurang
				nyawaku = nyawaku - 2
                fmt.Println("")
                if nyawa > 0 && nyawaku > 0 {
                    fmt.Println("Nyawa Kodok Acumalaka :", nyawa)
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Kodok Acumalaka use Backflip")
					fmt.Println("")
					fmt.Println("Nyawa Sigit Rendang :", nyawaku)
                } else if nyawa <= 0 && nyawaku > 0 || nyawa < nyawaku {
                    nyawa = 0
                    fmt.Println("Nyawa Kodok Acumalaka :", nyawa)

					
					fmt.Println("Kodok Acumalaka berhasil dikalahkan!")
					fmt.Println("")
					fmt.Println("Apakah kamu ingin menangkap Kodok Acumalaka?")
		
					fmt.Scan(&kodok)
					if kodok == "ya"{
						fmt.Println("")
						fmt.Println("Kamu memakai KepoBall")
						fmt.Println("")
						fmt.Println("Kodok Acumalaka berhasil ditangkap")
						fmt.Println("")
						fmt.Println("Kodok Acumalaka disimpan di tas")
					} else if kodok == "tidak" {
						fmt.Println("")
						fmt.Println("Okelah")
					}
		
					fmt.Println("")
					fmt.Println("Sekarang kamu bisa masuk ke Dunia Isekai")
					fmt.Println("")
					fmt.Println("Welcome!")
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Dibuat oleh Mikhael Restu Mahardhika")
					fmt.Println("103012300277")
					fmt.Println("Hehe TY ALL!")
                } else if nyawa > 0 && nyawaku <= 0 {
					fmt.Println("Kodok Acumalaka use Backflip")
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Kamu berhasil dikalahkan!")
					fmt.Println("Selamat tinggal!")
				} else if nyawa <= 0 && nyawaku <= 0 {
					fmt.Println("Kodok Acumalaka use Backflip")
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Kalian seri!!")
					fmt.Println("")
					fmt.Println("Kembalilah, bertarunglah lain waktu!")
				}

            }


        } else if y == "3" {
            sk1 = "Protein"
            sk2 = "Hasta la Vista"
            sk3 = "Brrrrrr"

            fmt.Println("")
            fmt.Println("Siap?")
            fmt.Println("")
            fmt.Println("Pertarungan dimulai HAHAHA")

            for nyawa > 0 && nyawaku > 0 {
                fmt.Println("")
                fmt.Println("")
                fmt.Println("Pilih Skill :")
                fmt.Println("1.", sk1)
                fmt.Println("2.", sk2)
                fmt.Println("3.", sk3)
                fmt.Scan(&z)
                fmt.Println("")

                if z == "1" {
                    ngurang = 1
                    fmt.Println("Soldier Use", sk1)
					fmt.Println("Ga Efektif, pake yang lain")
                } else if z == "2" {
                    ngurang = 2
                    fmt.Println("Soldier Use", sk2)
                } else if z == "3" {
                    ngurang = 3
                    fmt.Println("Soldier Use", sk3)
					fmt.Println("Efektif!!")
					fmt.Println("Musuh tutup mata!")
                }

                nyawa = nyawa - ngurang
				nyawaku = nyawaku - 2
                fmt.Println("")
                if nyawa > 0 && nyawaku > 0 {
                    fmt.Println("Nyawa Kodok Acumalaka :", nyawa)
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Kodok Acumalaka use Backflip")
					fmt.Println("")
					fmt.Println("Nyawa Soldier :", nyawaku)
                } else if nyawa <= 0 && nyawaku > 0 || nyawa < nyawaku {
                    nyawa = 0
                    fmt.Println("Nyawa Kodok Acumalaka :", nyawa)

					
					fmt.Println("Kodok Acumalaka berhasil dikalahkan!")
					fmt.Println("")
					fmt.Println("Apakah kamu ingin menangkap Kodok Acumalaka?")
		
					fmt.Scan(&kodok)
					if kodok == "ya"{
						fmt.Println("")
						fmt.Println("Kamu memakai KepoBall")
						fmt.Println("")
						fmt.Println("Kodok Acumalaka berhasil ditangkap")
						fmt.Println("")
						fmt.Println("Kodok Acumalaka disimpan di tas")
					} else if kodok == "tidak" {
						fmt.Println("")
						fmt.Println("Okelah")
					}
		
					fmt.Println("")
					fmt.Println("Sekarang kamu bisa masuk ke Dunia Isekai")
					fmt.Println("")
					fmt.Println("Welcome!")
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Dibuat oleh Mikhael Restu Mahardhika")
					fmt.Println("103012300277")
					fmt.Println("Hehe TY ALL!")
                } else if nyawa > 0 && nyawaku <= 0 {
					fmt.Println("Kodok Acumalaka use Backflip")
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Kamu berhasil dikalahkan!")
					fmt.Println("Selamat tinggal!")
				} else if nyawa <= 0 && nyawaku <= 0 {
					fmt.Println("Kodok Acumalaka use Backflip")
					fmt.Println("")
					fmt.Println("")
					fmt.Println("Kalian seri!!")
					fmt.Println("")
					fmt.Println("Kembalilah, bertarunglah lain waktu!")
				}

            }
            

        } else {
            fmt.Println("Cuma segitu gausah ngide!")
			fmt.Println("")
			fmt.Println("Udah balik aja sana!")
			fmt.Println("Kamu ga layak!")
        }


    } else if x == "tidak" {
        fmt.Println("Dih pengecut, pulang aja sana")
		fmt.Println("")
		fmt.Println("Apakah kamu ingin pulang?")
		fmt.Scan(&mau)

		if mau == "ya" {
			fmt.Println("Dih")

		} else if mau == "tidak" {
			fmt.Println("")
			fmt.Println("Hmmm, kudengar ada jalur belakang")
			fmt.Println("Kamu mau coba?")
			fmt.Scan(&mau2)

			if mau2 == "ya" {
				fmt.Println("")
				fmt.Println("Kamu berjalan menuju gerbang belakang")
				fmt.Println("")
				fmt.Println("Kamu bertemu Kakek Sugiono")
				fmt.Println("Untuk bisa masuk kamu harus membawa barang yang diminta Kakek Sugiono")
				fmt.Println("Kakek Sugiono meminta Miniatur Roket!!")
				fmt.Println("")
				fmt.Println("Apa kamu bersedia?")
				fmt.Scan(&mau3)

				if mau3 == "ya" {
					fmt.Println("")
					fmt.Println("Bagus, sekarang pergi ke Hutan Keramat!")
					fmt.Println("")
					fmt.Println("Kamu pergi ke Hutan Keramat")
					fmt.Println("")
					fmt.Println("Kamu bertemu Pocong Lari")
					fmt.Println("")
					fmt.Println("Apakah kamu mau lawan?")
					fmt.Println("lawan / kabur")
					fmt.Scan(&lawan)

					if lawan == "lawan" {
						fmt.Println("")
						fmt.Println("Ketiga petarungmu ketakutan!!!")
						fmt.Println("")
						fmt.Println("Kamu berhasil ditangkap Pocong Lari")
						fmt.Println("")
						fmt.Println("Kamu dibawa Pocong Lari ke Dunia Bawah")
						fmt.Println("")
						fmt.Println("Kasian")

					} else if lawan == "kabur" {
						fmt.Println("")
						fmt.Println("Kamu berhasil kabur dari Pocong Lari!")
						fmt.Println("")
						fmt.Println("Sekarang kamu menemukan 3 Miniatur Roket yang berbeda warna!")
						fmt.Println("Kamu harus pilih salah satu!")
						fmt.Println("")
						fmt.Println("Pilih miniatur roket : 1 / 2 / 3")
						fmt.Println("1. Merah")
						fmt.Println("2. Biru")
						fmt.Println("3. Pink Polkadot")
						fmt.Scan(&roket)

						if roket == "1" {
							benar = "tidak"
						} else if roket == "2" {
							benar = "tidak"
						} else if roket == "3" {
							benar = "ya"
						}

						fmt.Println("")
						fmt.Println("Kamu mengambil salah satu Miniatur Roket")
						fmt.Println("")
						fmt.Println("Kamu kembali ke Kakek Sugiono")
						fmt.Println("")
						fmt.Println("Kamu memberikan Miniatur Roket")

						if benar == "ya" {
							fmt.Println("")
							fmt.Println("Kakek Sugiono senang dengan Miniatur Roket yang kamu bawa!")
							fmt.Println("")
							fmt.Println("Kamu diperbolehkan masuk Dunia Isekai!")
							fmt.Println("")
							fmt.Println("Welcome")
							fmt.Println("")
							fmt.Println("")
							fmt.Println("Dibuat oleh Mikhael Restu Mahardhika")
							fmt.Println("103012300277")
							fmt.Println("HEHE TY ALL!!")

						} else if benar == "tidak" {
							fmt.Println("")
							fmt.Println("Kamu memberikan Miniatur Roket yang salah!!")
							fmt.Println("")
							fmt.Println("Kakek Sugiono murka!!!")
							fmt.Println("Kamu dikutuk menjadi Ular Sanca!!")
							fmt.Println("")
							fmt.Println("Kasian")

						}

					}

				} else if mau3 == "tidak" {
					fmt.Println("")
					fmt.Println("Aneh")
				}

			} else if mau2 == "tidak" {
				fmt.Println("")
				fmt.Println("Cih!!")

			}

		}
    }
    
}