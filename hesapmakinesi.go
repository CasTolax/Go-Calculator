package main

import "fmt"

func main() {
	var kullanıcı int
	var hesapsor int
	var sayı1 int
	var sayı2 int

	for {
		fmt.Println("hesap makinesini başlatmak için 1'e çıkmak için ise 2 ye basınız")
		fmt.Scanln(&kullanıcı)

		if kullanıcı == 1 {
			fmt.Println(" -- HESAP MAKİNESİ --")
			fmt.Println("1 = toplama 2 = çıkarma 3 = çarpma 4 = bolme")
			fmt.Println("lütfen seçim yapınız: ")
			fmt.Scanln(&hesapsor)

			if hesapsor == 1 {
				fmt.Println("1.sayı ve 2.sayı giriniz")
				fmt.Scanln(&sayı1)
				fmt.Scanln(&sayı2)

				hesapla := sayı1 + sayı2
				fmt.Println("sonuç:", hesapla)

			} else if hesapsor == 2 {
				fmt.Println("1.sayı ve 2.sayı giriniz")
				fmt.Scanln(&sayı1)
				fmt.Scanln(&sayı2)

				hesapla := sayı1 - sayı2
				fmt.Println("sonuç:", hesapla)

			} else if hesapsor == 3 {
				fmt.Println("1.sayı ve 2.sayı giriniz")
				fmt.Scanln(&sayı1)
				fmt.Scanln(&sayı2)

				hesapla := sayı1 * sayı2
				fmt.Println("sonuç:", hesapla)

			} else if hesapsor == 4 {
				fmt.Println("-- bölme işleminde 0 kullanmayınız! --")
				fmt.Println("1.sayı ve 2.sayı giriniz")

				fmt.Scanln(&sayı1)
				fmt.Scanln(&sayı2)

				hesapla := sayı1 / sayı2
				fmt.Println("sonuç:", hesapla)

			} else {
				fmt.Println("bilinmeyen hata")
				break
			}

		} else if kullanıcı == 2 {
			fmt.Println("program sonlandırıldı...")
			break

		} else {
			fmt.Println("hata:sadece 1'e veya 2'ye basabilirsiniz")
			break
		}

	}
}
