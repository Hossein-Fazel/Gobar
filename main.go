package main

import(
	"time"
	"github.com/Hossein-Fazel/Gobar/progressbar"
)

func main(){
	pbar := progressbar.NewProgressBar()
	pbar.Set_filled("-")  // Set filled character

	for i := 0; i <= 100; i++ {
		pbar.Update(1)
		pbar.Show()
		time.Sleep(50 * time.Millisecond)
	}

	pbar.Stop()



	pbar = progressbar.NewProgressBar()
	pbar.Set_total(100)
	pbar.Set_start(0)
	pbar.Set_size(30)
	pbar.Set_filled("■")  // Set filled character
	pbar.Set_emptyChar(" ") // Use a space as the empty character
	pbar.Set_displayMode("bar")

	for i := 0; i <= 100; i++ {
		pbar.Update(1)
		pbar.Show()
		time.Sleep(100 * time.Millisecond)
	}

	pbar.Stop()
}