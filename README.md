# GoBar - A Simple and Customizable Progress Bar in Go

**GoBar** is a lightweight and highly customizable progress bar library written in Go. It provides an easy way to visualize the progress of tasks in your command-line applications. Whether you're processing files, downloading data, or performing any other task that takes time, GoBar can help you keep track of progress in a clean and intuitive way.

## Features

- **Customizable Bar**: Set the filled character, empty character, total value, start value, and size of the progress bar.
- **Dynamic Updates**: Update the progress bar dynamically as your task progresses.
- **Reset Functionality**: Reset the progress bar to its initial state.
- **Stop Functionality**: Stop the progress bar and print a new line when the task is complete.
- **Unicode Support**: Use Unicode characters (e.g., `█`, `░`, `⬛`) for the filled and empty portions of the bar.
- **Multiple Display Formats**: Choose between full bar, percentage-only, or combined view

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

## Basic Usage

```go
package main

import (
	"time"
	"github.com/Hossein-Fazel/gobar/progressbar"
)

func main() {
	// Create a new progress bar
	pb := progressbar.NewProgressBar()
	pb.Set_total(100)
	
	// Update progress
	for i := 0; i <= 100; i++ {
		pb.Update(1)
		pb.Show()
		time.Sleep(50 * time.Millisecond)
	}
	pb.Stop()
}
```

## Display Modes Examples

### 1. Full Mode (default)
Shows bar with percentage

```go
pb := progressbar.NewProgressBar()
pb.Set_displayMode("full") // [##########] 50.0%
```

### 2. Bar Only Mode
Shows just the progress bar

```go
pb.Set_displayMode("bar") // [##########]
```

### 3. Percent Only Mode
Shows just the percentage

```go
pb.Set_displayMode("percent") // 50.0%
```

## Customization Examples

### Custom Characters
```go
pb := progressbar.NewProgressBar()
pb.Set_filled("■")
pb.Set_emptyChar("─")
pb.Set_displayMode("full") 
// [■■■■■■────] 60.0%
```

### Different Size
```go
pb := progressbar.NewProgressBar()
pb.Set_size(20) // Smaller bar
pb.Set_displayMode("full")
// [####      ] 40.0%
```

### Starting from a Position
```go
pb := progressbar.NewProgressBar()
pb.Set_start(30) // Start at 30%
pb.Set_displayMode("full")
// [######    ] 30.0%
```

## Advanced Example
```go
package main

import (
	"fmt"
	"time"
	"github.com/yourusername/gobar/progressbar"
)

func main() {
	fmt.Println("File Download Progress:")
	
	pb := progressbar.NewProgressBar()
	pb.Set_filled("=")
	pb.Set_emptyChar(" ")
	pb.Set_size(30)
	pb.Set_displayMode("full") // Try changing to "bar" or "percent"
	pb.Set_total(250)
	
	// Simulate download chunks
	for i := 0; i < 250; i += 10 {
		pb.Update(10)
		pb.Show()
		time.Sleep(100 * time.Millisecond)
	}
	
	pb.Stop()
	fmt.Println("Download complete!")
}
```

## API Reference

- `NewProgressBar()` - Creates new progress bar with defaults
- `Set_filled(string)` - Set filled character
- `Set_emptyChar(string)` - Set empty character
- `Set_total(int)` - Set total value
- `Set_start(int)` - Set starting value
- `Set_size(int)` - Set bar size
- `Set_displayMode(string)` - Set display mode ("full", "bar", "percent")
- `Update(int)` - Increment progress
- `Show()` - Display current progress
- `Stop()` - Complete progress bar
- `Reset()` - Reset progress bar

