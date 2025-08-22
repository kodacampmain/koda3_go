package main

import (
	"github.com/kodacampmain/koda3_go/internals/intermediate"
	m3 "github.com/kodacampmain/koda3_go/internals/minitask_3"
)

func main() {
	// utils.Greet()
	// // inference
	// // str := "Hello World"
	// // bt := []byte("golang")
	// // fmt.Printf("%s\n%v\n", str, bt)
	// // rn := []rune("ひ")
	// // rnj := []rune(" ꦈꦴ")
	// // fmt.Println(rn)
	// // fmt.Println(rnj)
	// // manifest
	// // var numb uint = 1
	// // fmt.Println(numb)
	// // angka menjadi hari
	// day, err := utils.NumToDay(0)
	// // explicit error handling
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(day)
	// }
	// total := numbers.Sum(1, 2, 3, 4, 5)
	// fmt.Println(total)
	// // for dengan range of int
	// // for i := range 3 {
	// // 	fmt.Println(i)
	// // }

	// collection.ArrayAndSlice()

	// collection.MapAndStruct()

	// batman := collection.Movies{
	// 	Title:  "Batman",
	// 	Genres: []string{"Action", "Romance"},
	// 	Id:     68,
	// }
	// fmt.Println(batman)
	// smk := minitask1.NewSchool("SMK Sumber Rejeki", "Electrical")
	// ithb := minitask1.NewSchool("Institut Teknologi Harapan Bangsa", "Electrical Engineering")
	// user := minitask1.NewPerson("Andi", "andi.webp", "andi@mail.com", "+6285768904329", 25, false, []minitask1.School{smk, ithb})
	// user := minitask1.Person{
	// 	Name:        "Andi",
	// 	Photo:       "andi.webp",
	// 	Email:       "andi@mail.com",
	// 	Age:         25,
	// 	PhoneNumber: "+6285768904329",
	// 	IsMarried:   false,
	// 	Schools: []minitask1.School{
	// 		{Name: "SMK Sumber Rezeki", Major: "Electrical"},
	// 		{Name: "Institut Teknologi Harapan Bangsa", Major: "Electrical Engineering"},
	// 	},
	// }
	// fmt.Println(user)
	// user.GetBiodata()
	// fmt.Printf("Sudah punya KTP: %t\n", user.IsEligibleForKTP())
	// user.UpdateName("Candra")
	// user.GetBiodata()

	// email := "a@mail.com"
	// pass := "pass"
	// success, err := minitask1.Login(email, pass)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Println(success)
	// }
	// fmt.Println(minitask1.BigNumber())

	// fromUnit := "C"
	// targetUnit := "F"
	// var startTemp float32 = 25
	// resultTemp, err := minitask1.TemperatureConversion(startTemp, fromUnit, targetUnit)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("%.2f%s -> %.2f%s\n", startTemp, fromUnit, resultTemp, targetUnit)
	// }
	// fmt.Println(minitask1.InsertIntoSlice([]int{50, 75, 66, 20, 32, 90}, 66, 88))
	// fmt.Println(minitask1.InsertIntoSlice([]string{"Hello", "World"}, "Hello", "Halo"))
	// defer fmt.Println("Program Selesai 1")
	// defer fmt.Println("Program Selesai 2")
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("panic recovered")
	// 		fmt.Println(r)
	// 	}
	// }()
	// intermediate.DeferredFunction()
	// intermediate.MustDo(false)
	// err := errors.New("terjadi error")
	// if err != nil {
	// 	intermediate.Exit(1, err)
	// }
	// fmt.Println("Proses")

	var day int = 20
	var ptrDay *int = &day

	*ptrDay = 10
	// fmt.Println(day)
	day = 30
	// fmt.Println(*ptrDay)
	// fmt.Println(&day, ptrDay)
	name := "Andi"
	// fmt.Println(name)
	intermediate.ChangeName(&name)
	// fmt.Println(name)
	// setter method

	// samsung := intermediate.TV{}
	// gree := intermediate.AC{}
	// miyako := intermediate.ElectricFan{}

	// intermediate.UniversalRemote("ON", &samsung)
	// intermediate.UniversalRemote("ON", &gree)
	// intermediate.UniversalRemote("ON", &miyako)
	// intermediate.UniversalRemote("OFF", &gree)
	// intermediate.UniversalRemote("OFF", &samsung)

	// result := service.MakePizza()
	// log.Printf("Pizza berhasil dibuat dengan keju %s dan saus tomat %s", result.Cheese.Type, result.Tomato.Brand)

	// mozarella := service.NewCheese("Mozarella")
	// cheddar := service.NewCheese("Cheddar")
	// heinz := service.NewTomato("Heinz")
	// ayam := service.NewTomato("Ayam")
	// betterResult1 := service.BetterMakePizza(mozarella, heinz)
	// betterResult2 := service.BetterMakePizza(mozarella, ayam)
	// betterResult3 := service.BetterMakePizza(cheddar, heinz)
	// betterResult4 := service.BetterMakePizza(cheddar, ayam)
	// service.LogPizza(betterResult1)
	// service.LogPizza(betterResult2)
	// service.LogPizza(betterResult3)
	// service.LogPizza(betterResult4)

	// if err := filereader.Read("./docs/file.txt"); err != nil {
	// 	log.Println(err.Error())
	// }
	// if err := filereader.Read("./docs/file"); err != nil {
	// 	log.Println(err.Error())
	// }
	// if err := filereader.Read("./docs"); err != nil {
	// 	log.Println(err.Error())
	// }

	// br := bank.NewBankRut("Bank Rut")
	// sp := online.NewSindbadPay("Sindbad Pay")
	// ttz := fiction.NewTTZ("TTZ")

	// checkout.Checkout(br, []int{10, 20, 30})
	// checkout.Checkout(sp, []int{30, 20, 40})
	// checkout.Checkout(ttz, []int{20, 20, 10})
	// checkout.Checkout(ttz, []int{40, 20, -15})
	// checkout.Checkout(ttz, []int{20, 50, 100})

	// log.Println(ttz.Payments)

	// fmt.Printf("Jumlah core yang available: %d\n", runtime.NumCPU())

	// var wg sync.WaitGroup
	// var mtx sync.Mutex
	// var rw sync.RWMutex
	// chn := make(chan string)

	// for range 4 {
	// 	wg.Add(3)
	// 	// go utils.Task(i, &wg)
	// 	// go utils.Task(i+10, &wg)
	// 	// wg.Add(1)
	// 	rw.RLock()
	// 	go utils.Read(&wg, chn, &rw, false)
	// 	go utils.Read(&wg, chn, &rw, true)
	// 	// out := <-chn
	// 	// fmt.Println(out)
	// 	// wg.Wait()
	// 	// wg.Add(1)
	// 	rw.Lock()
	// 	go utils.Write(&wg, chn, &rw)
	// 	// msg := <-chn
	// 	// log.Println(msg)
	// 	// wg.Wait()
	// }

	// time.Sleep(1 * time.Second)
	// wg.Wait() // blocking hingga wg counter bernilai 0

	// customers := []string{"Inkam", "Bonai", "Slamet", "Fepp", "Sdck"}
	// customerOrder := make(chan string, 2)
	// var wg sync.WaitGroup

	// wg.Add(2)
	// go service.Barista(customerOrder, 100200, &wg)
	// go service.Barista(customerOrder, 111222, &wg)
	// // Pemesanan
	// for _, customer := range customers {
	// 	fmt.Printf("%s places order\n", customer)
	// 	customerOrder <- fmt.Sprintf("%s's Latte", customer)

	// 	fmt.Printf("%s's order accepted (Queue size: %d)\n", customer, len(customerOrder))
	// 	time.Sleep(500 * time.Millisecond)
	// }

	// // Tutup toko
	// close(customerOrder)
	// // menutup channel harus dari sender
	// // time.Sleep(6 * time.Second)
	// wg.Wait()
	// fmt.Println("Coffee shop is closed")

	// orders := []struct {
	// 	Name string
	// 	Type string
	// }{{Name: "Roti", Type: "food"}, {Name: "Kopi", Type: "drink"}, {Name: "Nasi Goreng", Type: "food"}, {Name: "Teh", Type: "drink"}, {Name: "Mie Rebus", Type: "food"}, {Name: "Susu", Type: "drink"}}
	// makanan := make(chan string)
	// minuman := make(chan string)
	// selesai := make(chan bool)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go service.Kitchen(makanan, minuman, selesai, &wg)
	// for _, order := range orders {
	// 	if order.Type == "food" {
	// 		makanan <- order.Name
	// 		continue
	// 	}
	// 	if order.Type == "drink" {
	// 		minuman <- order.Name
	// 		continue
	// 	}
	// }
	// selesai <- true
	// wg.Wait()
	// fmt.Println("Kitchen Closed")
	// routines := []string{"Bangun", "Mandi", "Sarapan"}
	// var wg sync.WaitGroup

	// for i, routine := range routines {
	// 	wg.Add(1)
	// 	go minitask3.DoTask(routine, 100*(i+1), &wg)
	// 	wg.Wait()
	// }
	// // wg.Wait()
	// fmt.Println("Berangkat Kerja")
	familyChannel := make(chan m3.Message)
	go m3.NewWhiteBoard().Listen(familyChannel)

	m3.NewMessage("Mom", "Lets get breakfast").Send(familyChannel)
	m3.NewMessage("Dad", "I need help to fix the roof").Send(familyChannel)
	m3.NewMessage("Bigbro", "Come here, lemme smurf on your account").Send(familyChannel)
	m3.NewMessage("Brother 2", "Lets play FIFA").Send(familyChannel)
	m3.NewMessage("Sister", "Accompany me to a mall. Be my chaperone").Send(familyChannel)

	close(familyChannel)
}
