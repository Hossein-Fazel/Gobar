# GoBar - A Simple and Customizable Progress Bar in Go

**GoBar** is a lightweight and highly customizable progress bar library written in Go. It provides an easy way to visualize the progress of tasks in your command-line applications. Whether you're processing files, downloading data, or performing any other task that takes time, GoBar can help you keep track of progress in a clean and intuitive way.

## Features

- **Customizable Bar**: Set the filled character, empty character, total value, start value, and size of the progress bar.
- **Dynamic Updates**: Update the progress bar dynamically as your task progresses.
- **Reset Functionality**: Reset the progress bar to its initial state.
- **Stop Functionality**: Stop the progress bar and print a new line when the task is complete.
- **Unicode Support**: Use Unicode characters (e.g., `█`, `░`, `⬛`) for the filled and empty portions of the bar.
- **Percentage Display**: Show the progress percentage with one decimal place.

## Installation

To install GoBar, run the following command:

```bash
go get github.com/Hossein-Fazel/Gobar
```

To install a specific version (e.g., `v1.0.0`), run:

```bash
go get github.com/Hossein-Fazel/Gobar@v1.0.0
```

---

## Usage

### Importing the Package

To use GoBar in your Go project, import it as follows:

```go
import "github.com/Hossein-Fazel/Gobar"
```

### Creating a Progress Bar

To create a new progress bar, use the `NewProgressBar` function:

```go
pbar := progressbar.NewProgressBar()
```

### Customizing the Progress Bar

You can customize the progress bar by setting the filled character, empty character, total value, start value, and size:

```go
pbar.Set_filled("█")      // Set the filled character
pbar.Set_emptyChar("░")   // Set the empty character
pbar.Set_total(100)       // Set the total value
pbar.Set_start(0)         // Set the start value
pbar.Set_size(50)         // Set the size of the progress bar
```

### Updating the Progress Bar

To update the progress bar, use the `Update` method:

```go
pbar.Update(10)  // Update the progress bar by 10 units
```

### Resetting the Progress Bar

To reset the progress bar to its initial state, use the `Reset` method:

```go
pbar.Reset()
```

### Stopping the Progress Bar

To stop the progress bar and print a new line, use the `Stop` method:

```go
pbar.Stop()
```

### Displaying the Progress Bar

To display the progress bar, use the `Show` method:

```go
pbar.Show()
```

---

## Examples

### Example 1: Basic Usage

```go
package main

import (
	"time"
	"github.com/yourusername/gobar/progressbar"
)

func main() {
	pbar := progressbar.NewProgressBar()
	pbar.Set_filled("-")  // Set filled character

	pbar.Reset()

	for i := 0; i <= 100; i++ {
		pbar.Update(1)
		pbar.Show()
		time.Sleep(50 * time.Millisecond)
	}

	pbar.Stop()
}
```

---

### Example 2: Custom Filled and Empty Character

Use a custom filled and empty character to create a unique progress bar:

```go
package main

import (
	"time"
	"github.com/yourusername/gobar/progressbar"
)

func main() {
	pbar := progressbar.NewProgressBar()
	pbar.Set_total(100)
	pbar.Set_start(0)
	pbar.Set_size(30)
	pbar.Set_filled("■")  // Set filled character
	pbar.Set_emptyChar(" ") // Use a space as the empty character

	pbar.Reset()

	for i := 0; i <= 100; i++ {
		pbar.Update(1)
		pbar.Show()
		time.Sleep(100 * time.Millisecond)
	}

	pbar.Stop()
}
```